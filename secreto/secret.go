package secreto

const SecretKind = "Secret"
const ListKind = "List"

type SecretValue string

type KubeResource struct {
	Kind string `json:"kind"`
}

type DataMapper interface {
	DataMap() map[string]SecretValue
}

type Secret struct {
	// these fields are needed to validate the manifest
	Kind string `json:"kind"`
	// these are the only data that will be translated
	Data map[string]SecretValue `json:"data"`
}

func (s *Secret) DataMap() map[string]SecretValue {
	return s.Data
}

type List struct {
	// these fields are needed to validate the manifest
	Kind string `json:"kind"`

	// these are the only data that will be translated
	Items []*Secret `json:"items"`
}

func (l *List) DataMap() map[string]SecretValue {
	return nil
}
