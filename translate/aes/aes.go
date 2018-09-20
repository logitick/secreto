package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

type TextToAes struct {
	key []byte
}

type AesToText struct {
	Key []byte
}

func NewAesToTextTranslator(key []byte) *AesToText {
	decrypter := &AesToText{key}
	return decrypter
}

func NewAesTranslator(key string) (*TextToAes, error) {
	src := []byte(key)
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	src = append(src, padtext...)
	encrypter := &TextToAes{src}
	return encrypter, nil
}

func AesDecrypt(plaintext string, key string) string {
	d, _ := base64.StdEncoding.DecodeString(key)
	bkey, _ := hex.DecodeString(hex.EncodeToString([]byte(d)))
	decoder := TextToAes{bkey}
	return decoder.Translate(plaintext)
}

func (t *TextToAes) Translate(s string) string {
	block, err := aes.NewCipher(t.key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)

	copy(nonce, t.key[:12])
	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(s), nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func (d *AesToText) Translate(s string) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(s)
	nonce := make([]byte, 12)
	copy(nonce, d.Key[:12])
	block, err := aes.NewCipher(d.Key)
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
