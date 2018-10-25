package secreto

import (
	"fmt"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

func Out(res interface{}) error {
	return writer(res)
	// out, err := yaml.Marshal(res)
	// if err != nil {
	// 	return fmt.Errorf("Cannot marshal: %v", err)
	// }
	// oyaml := string(out)

	// fmt.Println(oyaml)
	// return nil
}

func writer(r interface{}) error {
	t := reflect.TypeOf(r)
	switch t {
	case reflect.TypeOf(&Secret{}):
		s, _ := r.(*Secret)
		return secretWriter(s)
	case reflect.TypeOf(&List{}):
		l, _ := r.(*List)
		return listWriter(l)
	}
	return fmt.Errorf("Missing writer")
}

func secretWriter(s *Secret) error {

	kv := make(map[string]interface{})
	yaml.Unmarshal(s.bytes, kv)
	kv["data"] = s.Data

	out, err := yaml.Marshal(kv)
	if err != nil {
		return fmt.Errorf("Cannot marshal: %v", err)
	}
	oyaml := string(out)

	fmt.Println(oyaml)
	return nil

}

func listWriter(l *List) error {
	kv := make(map[string]interface{})
	yaml.Unmarshal(l.bytes, kv)

	items := make([]interface{}, len(l.Items))

	for k, v := range kv["items"].([]interface{}) {
		skv := v.(map[interface{}]interface{})
		skv["data"] = l.Items[k].Data
		items[k] = skv
	}
	kv["items"] = items

	out, err := yaml.Marshal(kv)
	if err != nil {
		return fmt.Errorf("Cannot marshal: %v", err)
	}
	oyaml := string(out)

	fmt.Println(oyaml)
	return nil
}
