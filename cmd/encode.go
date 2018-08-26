package cmd

import (
	"os"

	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate/base64"

	"github.com/spf13/cobra"
)

var (
	encoder = new(base64.TextToBase64)
)

var cmdEncode = &cobra.Command{
	Use:   "encode [path to secrets.yml]",
	Short: "Encodes the literal values in a secrets file to base64", Long: `print is for printing anything back to the screen. For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reader, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		s, err := secreto.Read(reader)
		if err != nil {
			panic(err)
		}

		// data map - as in a map of the data values from the yaml
		dm := make(map[string]string)
		for k, v := range s.Data {
			dm[k] = encoder.Translate(v)
		}
		reader.Seek(0, 0)
		secreto.Write(reader, dm)
	},
}

func newEncodeCmd() *cobra.Command {
	return cmdEncode
}
