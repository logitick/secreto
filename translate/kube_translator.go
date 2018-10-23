package translate

import (
	"fmt"
	"reflect"

	"github.com/logitick/secreto/secreto"
)

type KubeTranslator interface {
	Translate(tt TextTranslator) interface{}
}

type MissingTranlatorError struct {
	resource interface{}
}

func (err *MissingTranlatorError) Error() string {
	return fmt.Sprintf("GetKubeTranslator: No translator for %v", err.resource)
}

func GetKubeTranslator(r interface{}) (KubeTranslator, error) {
	t := reflect.TypeOf(r)
	switch t {
	case reflect.TypeOf(&secreto.Secret{}):
		secret, _ := r.(*secreto.Secret)
		return &SecretTranslator{*secret}, nil
	case reflect.TypeOf(&secreto.List{}):
		l, _ := r.(*secreto.List)
		return &ListTranslator{*l}, nil
	}
	return nil, &MissingTranlatorError{r}
}

type SecretTranslator struct {
	Subject secreto.Secret
}

func (st *SecretTranslator) Translate(tt TextTranslator) interface{} {
	for k, v := range st.Subject.Data {
		st.Subject.Data[k] = secreto.SecretValue(tt.Translate(string(v)))
	}
	return &st.Subject
}

type ListTranslator struct {
	Subject secreto.List
}

func (lt *ListTranslator) Translate(tt TextTranslator) interface{} {
	for k, secret := range lt.Subject.Items {
		st, err := GetKubeTranslator(secret)
		if err != nil {
			panic(fmt.Errorf("ListTranslator: cannot translate secret: %v", err))
		}
		lt.Subject.Items[k] = st.Translate(tt).(*secreto.Secret)
	}
	return &lt.Subject
}
