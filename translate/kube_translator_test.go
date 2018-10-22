package translate

import (
	"reflect"
	"testing"
	s "github.com/logitick/secreto/secreto"
)

func TestGetKubeTranslator(t *testing.T) {
	tests := []struct {
		name string
		arg s.KubeResource
		want reflect.Type
	}{
		{name: "Secret type gets SecretTranslator", arg: s.KubeResource{Kind:"Secret"}, want: reflect.TypeOf(&SecretTranslator{})},
		{name: "Unkown type gets nil", arg: s.KubeResource{Kind:"Deployment"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKubeTranslator(tt.arg); reflect.TypeOf(got) != tt.want {
				t.Errorf("GetKubeTranslator() = got %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}
