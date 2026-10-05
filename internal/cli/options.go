package cli

import "runtime/debug"

// Option configures the root command [NewCommand] builds.
type Option func(*options)

type options struct {
	// readBuildInfo returns the build info the Go toolchain embeds in the
	// binary: [debug.ReadBuildInfo], unless a test replaces it.
	readBuildInfo func() (*debug.BuildInfo, bool)
	version       string
}

// WithVersion sets the release version `PROJECT_NAME version` reports. The
// commit, its time, the dirty state, the Go version and the platform come
// from the Go build info at run time. When version is empty, the main
// module's version from the build info is reported instead.
func WithVersion(version string) Option {
	return func(o *options) { o.version = version }
}

func defaultOptions() *options {
	return &options{readBuildInfo: debug.ReadBuildInfo}
}
