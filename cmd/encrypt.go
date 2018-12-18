package cmd

import (
	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate"

	"github.com/spf13/cobra"
)

var cmdEncrypt = &cobra.Command{
	Use:   "encrypt [path to secrets.yml]",
	Short: "Apply AES encryption to the values in a secrets file",
	Long:  `Apply AES encryption to the values in a secrets file. A key is required to apply the encryption`,
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

		tt, err := translate.NewAesTranslator(key)
		if err != nil {
			panic(err)
		}
		s = ktr.Translate(tt)
		w := secreto.GetWriter("")
		secreto.Out(s, w)
	},
}

func newEncryptCmd() *cobra.Command {
	cmdEncrypt.Flags().StringP("key", "k", "", "The key")
	cmdEncrypt.MarkFlagRequired("key")
	return cmdEncrypt
}
