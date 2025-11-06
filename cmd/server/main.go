package main

import (
	"fmt"
	"os"

	"github.com/OWNER/PROJECT_NAME/internal/cli/cmd"
)

func main() {
	cmdRoot := cmd.NewServerRootCmd()

	cmdRoot.AddCommand(serveCmd)

	err := cmdRoot.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
