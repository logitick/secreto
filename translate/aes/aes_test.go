package aes

import (
	"reflect"
	"testing"
)

func TestNewAesTranslator(t *testing.T) {
	tests := []struct {
		got  string
		want []byte
	}{
		{"secret_key", []byte{177, 231, 43}},
		{"$3cr3tK3y", []byte("$3cr3tK3y")},
	}

	for _, tt := range tests {
		t.Run(tt.got, func(t *testing.T) {

			if got := NewAesTranslator(tt.got); !reflect.DeepEqual(got.Key, tt.want) {
				t.Errorf("NewAesTranslator() = %v, want %v", got.Key, tt.want)
			}
		})
	}
}

func TestAesDecrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AesDecrypt(tt.args.plaintext, tt.args.key); got != tt.want {
				t.Errorf("AesDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
