package secreto

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

func Out(res interface{}) error {
	out, err := yaml.Marshal(res)
	if err != nil {
		return fmt.Errorf("Cannot marshal: %v", err)
	}
	oyaml := string(out)

	fmt.Println(oyaml)
	return nil
}
