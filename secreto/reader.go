package secreto

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func ReadFile(path string) ([]byte, error) {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: could not read bytes: %v", err)
	}
	return b, nil
}

func ReadBytes(b []byte, out interface{}) error {
	err := yaml.Unmarshal(b, out)
	if err != nil {
		return fmt.Errorf("ReadBytes: cannot Unmarshal: %v", err)
	}
	return nil
}

func GetResourceFromType(data []byte) (interface{}, *KubeResource, error) {
	kr := new(KubeResource)
	err := ReadBytes(data, kr)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot determine type: %v", err)
	}
	switch kr.Kind {
	case "Secret":
		s, err := ReadSecret(data)
		return s, kr, err
	case "List":
		l, err := ReadList(data)
		return l, kr, err
	}
	return nil, nil, fmt.Errorf("Cannot get resource of unknown type: %s", kr.Kind)
}

func ReadSecret(data []byte) (*Secret, error) {
	s := new(Secret)
	err := ReadBytes(data, s)
	if err != nil {
		return s, err
	}
	s.bytes = data
	err = ValidateSecret(s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func ValidateSecret(s *Secret) (err error) {
	if s.Kind != "Secret" {
		return errors.New("Invalid manifest kind. Expected secret got " + s.Kind)
	}
	return nil
}

func ReadList(data []byte) (*List, error) {
	l := new(List)
	err := ReadBytes(data, l)
	if err != nil {
		return l, err
	}
	l.bytes = data
	return l, err
}

func ValidateList(l *List) error {
	if l.Kind != "List" {
		return errors.New("Invalid manifest kind. Expected List got " + l.Kind)
	}
	return nil
}
