package secreto

import (
	"fmt"
	"io"
	"os"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

func GetWriter(writer string) io.Writer {
	var w io.Writer

	switch writer {
	default:
		w = os.Stdout
	}
	return w
}

func Out(res interface{}, w io.Writer) error {
	return writer(res, w)
}

func writer(r interface{}, w io.Writer) error {
	t := reflect.TypeOf(r)
	switch t {
	case reflect.TypeOf(&Secret{}):
		s, _ := r.(*Secret)
		return secretWriter(s, w)
	case reflect.TypeOf(&List{}):
		l, _ := r.(*List)
		return listWriter(l, w)
	}
	return fmt.Errorf("Missing writer")
}

func secretWriter(s *Secret, w io.Writer) error {

	kv := make(map[string]interface{})
	yaml.Unmarshal(s.bytes, kv)
	kv["data"] = s.Data

	out, err := yaml.Marshal(kv)
	if err != nil {
		return fmt.Errorf("Cannot marshal: %v", err)
	}
	_, err = w.Write(out)
	return err
}

func listWriter(l *List, w io.Writer) error {
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
	_, err = w.Write(out)
	return err
}
