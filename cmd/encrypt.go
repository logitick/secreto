package cmd

import (
	"os"

	"github.com/logitick/secreto/secreto"
	"github.com/logitick/secreto/translate/aes"

	"github.com/spf13/cobra"
)

var (
	encrypterLongDesc = ``
)

var cmdEncrypt = &cobra.Command{
	Use:   "encrypt [path to secrets.yml]",
	Short: "",
	Long:  encrypterLongDesc,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reader, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		key, err := cmd.Flags().GetString("key")
		if err != nil {
			panic(err)
		}

		//bSliceKey, _ := hex.DecodeString(key)
		encrypter, err := aes.NewAesTranslator(key)
		if err != nil {
			panic(err)
		}

		s, err := secreto.Read(reader)
		if err != nil {
			panic(err)
		}

		// data map - as in a map of the data values from the yaml
		dm := make(map[string]string)
		for k, v := range s.Data {
			dm[k] = encrypter.Translate(v)
		}
		reader.Seek(0, 0)
		secreto.Write(reader, dm)
	},
}

func newEncryptCmd() *cobra.Command {
	cmdEncrypt.Flags().StringP("key", "k", "", "The string key used to encrypt the values e.g. 2c53acf39d14a797c8822ac4d2dc9153597eca74a828addd846255673817f513 ")
	cmdEncrypt.MarkFlagRequired("key")
	return cmdEncrypt
}
