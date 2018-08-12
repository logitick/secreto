package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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

		s, b, err := secreto.Read(reader)
		if err != nil {
			panic(err)
		}

		dm := make(map[string]interface{}) // data map
		for k, v := range s.Data {
			dm[k] = encoder.Translate(v)
		}

		// create a map represenation of the secrets file
		// because secreto.Secrets only holds the data needed
		// to be validated and translated. If that were to
		// be unmarshalled to yaml, the user would lose the other
		// properties defined in the manifest that isn't in
		// the Secrets struct
		m := make(map[string]interface{})
		err = secreto.ReadBytes(b, m)
		if err != nil {
			panic(err)
		}
		m["data"] = dm

		out, err := yaml.Marshal(m)
		if err != nil {
			panic(err)
		}
		oyaml := string(out)

		fmt.Println(oyaml)
	},
}

func newEncodeCmd() *cobra.Command {
	cmdEncode.Flags().StringP("file", "f", "", "the path to the secrets manifest")
	cmdEncode.MarkFlagRequired("file")
	return cmdEncode
}
