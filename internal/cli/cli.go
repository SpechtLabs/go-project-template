// Package cli builds the PROJECT_NAME command tree. [NewCommand] builds the
// root command with every subcommand attached, and [Execute] runs it with
// styled help and errors and turns the outcome into an exit status.
//
// Add a subcommand by writing a constructor like newVersionCommand that takes
// what it needs as arguments, and attaching it in NewCommand. Commands return
// humane errors (github.com/sierrasoftworks/humane-errors-go), which carry
// advice on what to do next along with what went wrong.
package cli

import (
	"context"
	"os"
	"syscall"

	"charm.land/fang/v2"
	"github.com/spf13/cobra"
)

// NewCommand returns the PROJECT_NAME root command with every subcommand
// attached, configured by opts. Each call builds a new command tree.
func NewCommand(opts ...Option) *cobra.Command {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	cmd := &cobra.Command{
		Use:   "PROJECT_NAME",
		Short: "PROJECT_NAME does one thing, and does it well",
		Long: `PROJECT_NAME is a placeholder description. Say here what the tool is for,
in a sentence or two, and leave the details to the subcommands.`,
		// The error handler prints errors, and a usage dump after every error
		// buries the message.
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(newVersionCommand(*o))
	return cmd
}

// Execute runs cmd with args, styling help, usage and errors, and returns
// the process's exit status: 1 when the command failed, after printing why,
// and 0 otherwise. An interrupt or SIGTERM cancels the context the command
// runs with.
func Execute(ctx context.Context, cmd *cobra.Command, args []string) int {
	cmd.SetArgs(args)

	// `PROJECT_NAME version` is the single source of version information, so
	// fang's --version flag is off.
	err := fang.Execute(ctx, cmd,
		fang.WithoutVersion(),
		fang.WithNotifySignal(os.Interrupt, syscall.SIGTERM),
	)
	if err != nil {
		return 1
	}
	return 0
}
