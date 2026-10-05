package cli

import (
	"bytes"
	"context"
	"errors"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStatus int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "version",
			args:       []string{"version"},
			wantStatus: 0,
			wantStdout: "Version:     1.2.3\n",
		},
		{
			name:       "version rejects arguments",
			args:       []string{"version", "extra"},
			wantStatus: 1,
			wantStderr: `Unknown command "extra"`,
		},
		{
			name:       "unknown command",
			args:       []string{"frobnicate"},
			wantStatus: 1,
			wantStderr: `Unknown command "frobnicate"`,
		},
		{
			name:       "help",
			args:       []string{"--help"},
			wantStatus: 0,
			wantStdout: "version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(WithVersion("1.2.3"))
			var stdout, stderr bytes.Buffer
			cmd.SetOut(&stdout)
			cmd.SetErr(&stderr)

			if got := Execute(context.Background(), cmd, tt.args); got != tt.wantStatus {
				t.Errorf("Execute(%q) = %d, want %d; stderr:\n%s", tt.args, got, tt.wantStatus, stderr.String())
			}
			if !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, want it to contain %q", stdout.String(), tt.wantStdout)
			}
			if !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr = %q, want it to contain %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}

func TestNewBuildInfo(t *testing.T) {
	platform := runtime.GOOS + "/" + runtime.GOARCH

	tests := []struct {
		buildInfo *debug.BuildInfo
		name      string
		version   string
		want      buildInfo
	}{
		{
			name:    "release build",
			version: "1.2.3",
			buildInfo: &debug.BuildInfo{
				Main: debug.Module{Version: "v1.2.3"},
				Settings: []debug.BuildSetting{
					{Key: "vcs.revision", Value: "abc123"},
					{Key: "vcs.time", Value: "2026-01-02T03:04:05Z"},
					{Key: "vcs.modified", Value: "false"},
				},
			},
			want: buildInfo{Version: "1.2.3", Commit: "abc123", CommitTime: "2026-01-02T03:04:05Z", GoVersion: runtime.Version(), Platform: platform},
		},
		{
			name:    "local build of a modified tree",
			version: "",
			buildInfo: &debug.BuildInfo{
				Main: debug.Module{Version: "(devel)"},
				Settings: []debug.BuildSetting{
					{Key: "vcs.revision", Value: "abc123"},
					{Key: "vcs.modified", Value: "true"},
				},
			},
			want: buildInfo{Version: "(devel)", Commit: "abc123", CommitTime: unknown, GoVersion: runtime.Version(), Platform: platform, Dirty: true},
		},
		{
			name:      "no build info",
			version:   "",
			buildInfo: nil,
			want:      buildInfo{Version: unknown, Commit: unknown, CommitTime: unknown, GoVersion: runtime.Version(), Platform: platform},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := defaultOptions()
			WithVersion(tt.version)(o)
			o.readBuildInfo = func() (*debug.BuildInfo, bool) { return tt.buildInfo, tt.buildInfo != nil }

			if got := newBuildInfo(*o); got != tt.want {
				t.Errorf("newBuildInfo() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPrintVersion(t *testing.T) {
	tests := []struct {
		name    string
		info    buildInfo
		want    string
		wantErr bool
	}{
		{
			name: "clean tree",
			info: buildInfo{Version: "1.2.3", Commit: "abc123", CommitTime: "t", GoVersion: "go", Platform: "p"},
			want: "Version:     1.2.3\nCommit:      abc123\nCommit time: t\nGo version:  go\nPlatform:    p\n",
		},
		{
			name: "dirty tree",
			info: buildInfo{Version: "1.2.3", Commit: "abc123", CommitTime: "t", GoVersion: "go", Platform: "p", Dirty: true},
			want: "Version:     1.2.3\nCommit:      abc123 (dirty)\nCommit time: t\nGo version:  go\nPlatform:    p\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := printVersion(&out, tt.info); err != nil {
				t.Fatalf("printVersion() error = %v", err)
			}
			if out.String() != tt.want {
				t.Errorf("printVersion() wrote %q, want %q", out.String(), tt.want)
			}
		})
	}
}

func TestPrintVersionWriteError(t *testing.T) {
	err := printVersion(failingWriter{}, buildInfo{})
	if err == nil {
		t.Fatal("printVersion() to a failing writer returned no error")
	}
	if !errors.Is(err, errWrite) {
		t.Errorf("printVersion() error = %v, want it to wrap %v", err, errWrite)
	}
}

var errWrite = errors.New("write failed")

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errWrite }
