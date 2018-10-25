package translate

import (
	"testing"

	s "github.com/logitick/secreto/secreto"
	"github.com/stretchr/testify/assert"
)

func TestGetKubeTranslator(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want KubeTranslator
		err  error
	}{
		{name: "Secret type gets SecretTranslator and nil error", arg: new(s.Secret), want: &SecretTranslator{}, err: nil},
		{name: "List type gets ListTranslator and nil error", arg: new(s.List), want: &ListTranslator{}, err: nil},
		{name: "Unkown type gets nil with error", arg: s.KubeResource{Kind: "Deployment"}, want: nil, err: &MissingTranlatorError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKubeTranslator(tt.arg)
			assert.IsType(t, tt.want, got)
			assert.IsType(t, tt.err, err)
		})
	}
}
