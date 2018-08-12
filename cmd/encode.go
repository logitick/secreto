package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

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
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(errors.New("cannot read file: " + path))
		}
		s := new(secreto.Secret)
		err = yaml.Unmarshal(bytes, s)
		if err != nil {
			panic(errors.New("Cannot parse yaml: " + path))
		}
		if s.Version != "v1" {
			panic(errors.New("Invalid manifest version. Expected v2 got " + s.Version))
		}
		if s.Kind != "Secret" {
			panic(errors.New("Invalid manifest kind. Expected secret got " + s.Kind))
		}

		m := make(map[string]interface{})
		err = yaml.Unmarshal(bytes, m)

		dm := make(map[string]interface{}) // data map
		for k, v := range s.Data {
			dm[k] = encoder.Translate(v)
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
