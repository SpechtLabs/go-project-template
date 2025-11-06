package main

import (
	"fmt"
	"os"

	"github.com/OWNER/PROJECT_NAME/internal/cli/cmd"
)

func main() {
	cmdRoot := cmd.NewCliRootCmd()

	// TODO: Add your commands here
	// Example:
	// cmdRoot.AddCommand(cmdYourCommand)

	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
