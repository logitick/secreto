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
	Short: "Encodes the literal values in a secrets file to base64", Long: `print is for printing anything back to the screen. For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b, err := secreto.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		s, _, err := secreto.GetResourceFromType(b)
		if err != nil {
			panic(err)
		}
		err = secreto.ReadBytes(b, s)
		if err != nil {
			panic(err)
		}

		ttr := new(translate.TextToBase64)

		ktr, err := translate.GetKubeTranslator(s)
		if err != nil {
			panic(err)
		}
		s = ktr.Translate(ttr)
		secreto.Out(s)
	},
}

func newEncodeCmd() *cobra.Command {
	return cmdEncode
}
