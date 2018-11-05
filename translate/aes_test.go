package translate

import (
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

func TestTextToAes_Translate(t *testing.T) {
	tests := []struct {
		key       []byte
		plaintext string
		want      string
	}{
		{[]byte("helloworld123456"), "nonono", "KCR3q7BKOlg/8Feo:yEwJOOlKhvKSvbJ44FpISVKGErPkPQ=="},
		{[]byte("secret_key$ecure"), "nonono", "KCR3q7BKOlg/8Feo:cYz9VSWCIjzmd0Ci1CZOcxxStpcjiw=="},
		{[]byte("secret_key$ecure"), "hello there", "KCR3q7BKOlg/8Feo:d4b/ViTNo7gBEUMjLL9YLO/v331Ji3hKKvdY"},
	}
	for _, tt := range tests {
		t.Run(tt.plaintext, func(t *testing.T) {
			tta := &TextToAes{
				key:   tt.key,
				nonce: new(randomNonce),
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
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &AesToText{
				Key: tt.fields.Key,
			}
			if got := d.Translate(tt.args.s); got != tt.want {
				t.Errorf("AesToText.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
