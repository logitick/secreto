package cmd

import (
	"errors"
	"os"

	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate"

	"github.com/spf13/cobra"
)

var (
	encoder = new(translate.Base64)
)

var cmdEncode = &cobra.Command{
	Use:   "encode [path to secrets.yml]",
	Short: "Encodes the literal values in a secrets file to base64", Long: `print is for printing anything back to the screen. For many years people have printed back to the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(errors.New("Path is not specified"))
		}

		reader, err := os.Open(path)
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
	cmdEncode.Flags().StringP("file", "f", "", "the path to the secrets manifest")
	cmdEncode.MarkFlagRequired("file")
	return cmdEncode
}
