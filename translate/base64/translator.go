package base64

import "encoding/base64"

type TextToBase64 struct{}

func (b64 *TextToBase64) Translate(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

type Base64ToText struct{}

func (t *Base64ToText) Translate(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(decoded)
}
