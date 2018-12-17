package translate

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestNewAesTranslator(t *testing.T) {
	tests := []struct {
		plaintextKey string
		keyBytes     []byte
	}{
		{"secret_key", []byte("secret_key")},
		{"$3cr3tK3y", []byte("$3cr3tK3y")},
	}

	for _, tt := range tests {
		t.Run(tt.plaintextKey, func(t *testing.T) {
			if got, err := NewAesTranslator(tt.plaintextKey); err != nil && !reflect.DeepEqual(got.key, tt.keyBytes) {
				t.Errorf("Key = %v, want %v", got.key, tt.keyBytes)
			}
		})
	}
}

type testNonceGenerator struct {
	n []byte
}

func (tng testNonceGenerator) nonce() []byte {
	return tng.n
}

func TestTextToAes_Translate(t *testing.T) {

	tests := []struct {
		key       []byte
		plaintext string
		want      string
	}{
		{[]byte("helloworld123456"), "nonono", "AAAAAAAAAAAAAAAA:QRM8ozUKp+jif3e5qpZpiRgQXB0CEw=="},
		{[]byte("secret_key$ecure"), "nonono", "AAAAAAAAAAAAAAAA:RcU6zzwehJosI+JHh6DVDm1BMG49jw=="},
		{[]byte("secret_key$ecure"), "hello there", "AAAAAAAAAAAAAAAA:Q884zD1RmmWyjt+faMMoQZJdOlqkqD/oo+pc"},
	}

	for _, tt := range tests {
		t.Run(tt.plaintext, func(t *testing.T) {
			tta := &TextToAes{
				key:   tt.key,
				nonce: testNonceGenerator{n: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			}
			if got := tta.Translate(tt.plaintext); got != tt.want {
				t.Errorf("TextToAes.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesToText_Translate(t *testing.T) {
	type fields struct {
		Key []byte
	}
	type args struct {
		s string
	}
	tests := []struct {
		encryptedString string
		key             []byte
		want            string
	}{
		{"AAAAAAAAAAAAAAAA:QRM8ozUKp+jif3e5qpZpiRgQXB0CEw==", []byte("helloworld123456"), "nonono"},
		{"AAAAAAAAAAAAAAAA:RcU6zzwehJosI+JHh6DVDm1BMG49jw==", []byte("secret_key$ecure"), "nonono"},
		{"AAAAAAAAAAAAAAAA:Q884zD1RmmWyjt+faMMoQZJdOlqkqD/oo+pc", []byte("secret_key$ecure"), "hello there"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			d := &AesToText{
				key: tt.key,
			}
			if got := d.Translate(tt.encryptedString); got != tt.want {
				t.Errorf("AesToText.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAesToTextTranslator_decodes_base64_keys(t *testing.T) {
	token := make([]byte, 48)
	rand.Read(token)
	s := fmt.Sprintf("%X", token)

	fmt.Printf("%s", s)

	b := base64.StdEncoding.EncodeToString(token)
	fmt.Printf("%s", b)

	tests := []struct {
		key         string
		bytesString string
		byteLength  int
	}{
		{"base64:Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9HixkmBhVrYaB0NhtHpHgAWeTnL", "52fdfc072182654f163f5f0f9a621d729566c74d10037c4d7bbb0407d1e2c64981855ad8681d0d86d1e91e00167939cb", 48},
		{"base64:Uv38ByGCZU8WP18PmmIdcg==", "52fdfc072182654f163f5f0f9a621d72", 16},
		{"base64:Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9Hixkk=", "52fdfc072182654f163f5f0f9a621d729566c74d10037c4d7bbb0407d1e2c649", 32},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			want, _ := hex.DecodeString(tt.bytesString)
			if got := NewAesToTextTranslator(tt.key); !reflect.DeepEqual(got.key[:tt.byteLength], want) { // chop of everything after the key length because they're only padding
				t.Errorf("NewAesToTextTranslator() = %v, want %v", got.key, want)
			}
		})
	}
}
