package translate

import "testing"

func TestTextToBase64_Translate(t *testing.T) {
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
			b64 := &TextToBase64{}
			if got := b64.Translate(tt.subject); got != tt.want {
				t.Errorf("Base64.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64ToString_Translate(t *testing.T) {
	tests := []struct {
		subject string
		want    string
	}{
		{"aHVudGVyMg==", "hunter2"},
		{"cGFycm90LmxpdmU=", "parrot.live"},
		{"c3VwZXJjYWxpZnJhZ2lsaXN0aWM=", "supercalifragilistic"},
	}
	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {
			b64 := &Base64ToText{}
			if got := b64.Translate(tt.subject); got != tt.want {
				t.Errorf("Base64.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
