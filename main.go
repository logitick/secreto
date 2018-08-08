package main

import (
	"github.com/logitick/secreto/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "secreto"}
	for _, cmd := range cmd.Commands {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()
}
