package main

import (
	"context"
	"fmt"
	"os"

	"github.com/OWNER/PROJECT_NAME/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serveCmd.PersistentFlags().Int("port", 8080, "Port for the server")
	viper.SetDefault("server.port", 8080)
	if err := viper.BindPFlag("server.port", serveCmd.PersistentFlags().Lookup("port")); err != nil {
		panic(fmt.Errorf("fatal binding flag: %w", err))
	}
}

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Run the PROJECT_NAME server",
		Long:  `Start the PROJECT_NAME server.`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runE(cmd, args); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func runE(cmd *cobra.Command, _ []string) error {
	debug := viper.GetBool("debug")
	port := viper.GetInt("server.port")

	ctx, cancelFn := context.WithCancelCause(cmd.Context())
	defer cancelFn(context.Canceled)

	utils.InterruptHandler(ctx, cancelFn)

	// TODO: Implement your server logic here
	fmt.Printf("Starting server on port %d (debug: %v)\n", port, debug)

	// Wait for context done
	<-ctx.Done()

	fmt.Println("Server shutting down...")
	return nil
}
