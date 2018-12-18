package cmd

import (
	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate"

	"github.com/spf13/cobra"
)

var (
	decoder = new(translate.Base64ToText)
)

var cmdDecode = &cobra.Command{
	Use:   "decode [path to secrets.yml]",
	Short: "Translate the base64 encoded values to plain text",
	Long:  `Decodes your base64 secrets values into readable plain text values`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		s = ktr.Translate(decoder)
		w := secreto.GetWriter("")
		secreto.Out(s, w)
	},
}

func newDecodeCmd() *cobra.Command {
	return cmdDecode
}
