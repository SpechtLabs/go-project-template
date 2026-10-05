// Command PROJECT_NAME is the PROJECT_NAME command line. The command tree
// lives in internal/cli, where it can be tested; this is only the entry point.
package main

import (
	"context"
	"os"

	"github.com/spechtlabs/PROJECT_NAME/internal/cli"
)

// version is set via ldflags -X at release time (see .goreleaser.yaml). Commit,
// commit time and dirty state come from the VCS info the Go toolchain embeds
// (see `PROJECT_NAME version`).
var version string

func main() {
	os.Exit(cli.Execute(context.Background(), cli.NewCommand(cli.WithVersion(version)), os.Args[1:]))
}
