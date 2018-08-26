package translate

type Translator interface {
	Translate(s string) string
}
