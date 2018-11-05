package cmd

import (
	"github.com/spf13/cobra"
)

const version = "0.1.0"

// Commands contains the available cobra commands
var Commands []*cobra.Command

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "show the current version",
	Long:  `Display the version number of your secreto installation`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("0.0.1")
		return
	},
}

func init() {
	Commands = append(Commands,
		cmdVersion,
		newEncodeCmd(),
		newDecodeCmd(),
		newEncryptCmd(),
		newDecryptCmd(),
	)
}
