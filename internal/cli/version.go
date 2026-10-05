package cli

import (
	"fmt"
	"io"
	"runtime"

	"github.com/sierrasoftworks/humane-errors-go"
	"github.com/spf13/cobra"
)

// unknown stands in for build details the binary doesn't carry, such as the
// commit of a binary built with go run.
const unknown = "unknown"

// buildInfo is what `PROJECT_NAME version` reports.
type buildInfo struct {
	Version    string
	Commit     string
	CommitTime string
	GoVersion  string
	Platform   string
	Dirty      bool
}

// newVersionCommand returns the version command.
func newVersionCommand(o options) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version and build information",
		Long: `Shows the release version, plus the commit, commit time, Go version and
platform the binary was built with.

The commit details come from the version control information the Go toolchain
embeds in every build. Binaries built with go run don't carry it, so they report
those fields as unknown.`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return printVersion(cmd.OutOrStdout(), newBuildInfo(o))
		},
	}
}

// newBuildInfo combines the injected version with the VCS and toolchain
// details the Go toolchain embeds in the binary.
func newBuildInfo(o options) buildInfo {
	info := buildInfo{
		Version:    o.version,
		Commit:     unknown,
		CommitTime: unknown,
		GoVersion:  runtime.Version(),
		Platform:   runtime.GOOS + "/" + runtime.GOARCH,
	}

	if bi, ok := o.readBuildInfo(); ok {
		if info.Version == "" {
			info.Version = bi.Main.Version
		}
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				info.Commit = s.Value
			case "vcs.time":
				info.CommitTime = s.Value
			case "vcs.modified":
				info.Dirty = s.Value == "true"
			}
		}
	}

	if info.Version == "" {
		info.Version = unknown
	}
	return info
}

// printVersion writes info to w, one detail per line.
func printVersion(w io.Writer, info buildInfo) humane.Error {
	commit := info.Commit
	if info.Dirty {
		commit += " (dirty)"
	}

	_, err := fmt.Fprintf(w, "Version:     %s\nCommit:      %s\nCommit time: %s\nGo version:  %s\nPlatform:    %s\n",
		info.Version, commit, info.CommitTime, info.GoVersion, info.Platform)
	if err != nil {
		return humane.Wrap(err, "failed to print the version information", "check that standard output is writable, e.g. not a closed pipe")
	}
	return nil
}
