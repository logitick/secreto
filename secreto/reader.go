package secreto

import (
	"errors"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func ValidateSecret(s *Secret) (err error) {
	if s.Version != "v1" {
		return errors.New("Invalid manifest version. Expected v2 got " + s.Version)
	}
	if s.Kind != "Secret" {
		return errors.New("Invalid manifest kind. Expected secret got " + s.Kind)
	}
	return nil
}

func Read(r io.Reader) (*Secret, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s, err := ReadFromBytes(b)
	return s, err
}

func ReadFromBytes(b []byte) (*Secret, error) {
	s := new(Secret)
	err := yaml.Unmarshal(b, s)
	if err != nil {
		return s, err
	}
	err = ValidateSecret(s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func ReadBytes(b []byte, out interface{}) error {
	err := yaml.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}
