package cmd

import (
	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate"

	"github.com/spf13/cobra"
)

var cmdDecrypt = &cobra.Command{
	Use:   "decrypt [path to secrets.yml]",
	Short: "Reverses the AES encryption back to its original value.",
	Long:  `Decrypts the AES encrypted secrets.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			panic(err)
		}

		b, err := secreto.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		s, _, err := secreto.GetResourceFromType(b)
		if err != nil {
			panic(err)
		}

		ktr, err := translate.GetKubeTranslator(s)
		if err != nil {
			panic(err)
		}
		tt := translate.NewAesToTextTranslator(key)
		s = ktr.Translate(tt)
		w := secreto.GetWriter("")
		secreto.Out(s, w)
	},
}

func newDecryptCmd() *cobra.Command {
	cmdDecrypt.Flags().StringP("key", "k", "", "The key")
	cmdDecrypt.MarkFlagRequired("key")
	return cmdDecrypt
}
