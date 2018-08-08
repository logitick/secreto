package secreto

type Translator interface {
	Translate(s string) string
}
