package translate

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
)

type TextToAes struct {
	key   []byte
	nonce NonceGenerator
}

type AesToText struct {
	key []byte
}

type NonceGenerator interface {
	nonce() []byte
}

type randomNonce struct{}

func (ng *randomNonce) nonce() []byte {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	return nonce
}

func base64DecodeKey(key string) []byte {
	bkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return bkey
}

func NewAesToTextTranslator(key string) *AesToText {
	src := []byte(key)
	if strings.HasPrefix(key, "base64:") {
		src = base64DecodeKey(key[7:]) // strip "base64:"
	}
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	src = append(src, padtext...)
	decrypter := &AesToText{src}
	return decrypter
}

func NewAesTranslator(key string) (*TextToAes, error) {
	src := []byte(key)
	if strings.HasPrefix(key, "base64:") {
		src = base64DecodeKey(key[7:]) // strip "base64:"
	}
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	src = append(src, padtext...)
	encrypter := &TextToAes{src, new(randomNonce)}
	return encrypter, nil
}

func (t *TextToAes) Translate(s string) string {
	block, err := aes.NewCipher(t.key)
	if err != nil {
		panic(err.Error())
	}

	nonce := t.nonce.nonce()
	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(s), nil)
	return string(base64.StdEncoding.EncodeToString(nonce)) + ":" + string(base64.StdEncoding.EncodeToString(ciphertext))
}

func (d *AesToText) Translate(s string) string {
	in := strings.Split(s, ":")
	ciphertext, _ := base64.StdEncoding.DecodeString(in[1])

	nonce, _ := base64.StdEncoding.DecodeString(in[0])
	block, err := aes.NewCipher(d.key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
