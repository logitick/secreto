package cmd

import (
	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate"

	"github.com/spf13/cobra"
)

var (
	encoder = new(translate.TextToBase64)
)

var cmdEncode = &cobra.Command{
	Use:   "encode [path to secrets.yml]",
	Short: "Encodes the literal values in a secrets file to base64",
	Long:  ``,
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
		s = ktr.Translate(encoder)
		w := secreto.GetWriter("")
		secreto.Out(s, w)
	},
}

func newEncodeCmd() *cobra.Command {
	return cmdEncode
}
