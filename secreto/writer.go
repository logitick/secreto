package secreto

import (
	"fmt"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Write prints the output to stdout
func Write(r io.Reader, translated map[string]string) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// create a map represenation of the secrets file
	// because secreto.Secrets only holds the data needed
	// to be validated and translated. If that were to
	// be unmarshalled to yaml, the user would lose the other
	// properties defined in the manifest that isn't in
	// the Secrets struct
	m := make(map[string]interface{})
	err = ReadBytes(b, m)
	if err != nil {
		panic(err)
	}
	m["data"] = translated

	out, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	oyaml := string(out)

	fmt.Println(oyaml)

	return nil
}
