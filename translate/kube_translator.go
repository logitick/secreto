package translate

import "github.com/logitick/secreto/secreto"

type KubeTranslator interface {
	Translate(tt TextTranslator) *secreto.Secret
}

func GetKubeTranslator(res secreto.KubeResource) KubeTranslator {

	switch res.Kind {
	case secreto.SecretKind:
		return new(SecretTranslator)
	}
	return nil
}

type SecretTranslator struct {
	Subject secreto.Secret
}

func (st *SecretTranslator) Translate(tt TextTranslator) *secreto.Secret {
	for k, v := range st.Subject.Data {
		st.Subject.Data[k] = secreto.SecretValue(tt.Translate(string(v)))
	}
	return &st.Subject
}
