package translate

import "encoding/base64"

type Translator interface {
	Translate(s string) string
}

type Base64 struct{}

func (b64 *Base64) Translate(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
