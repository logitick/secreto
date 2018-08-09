package translate

import "testing"

func TestBase64_Translate(t *testing.T) {
	tests := []struct {
		subject string
		want    string
	}{
		{"hello", "aGVsbG8="},
		{"world", "d29ybGQ="},
		{"supercalifragilistic", "c3VwZXJjYWxpZnJhZ2lsaXN0aWM="},
		{"golang", "Z29sYW5n"},
	}
	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {
			b64 := &Base64{}
			if got := b64.Translate(tt.subject); got != tt.want {
				t.Errorf("Base64.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
