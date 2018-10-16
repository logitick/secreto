package secreto

type SecretValue string

type KubeResource struct {
	Kind string `json:"kind"`
}

type Secret struct {
	// these fields are needed to validate the manifest
	Kind string `json:"kind"`
	// these are the only data that will be translated
	Data map[string]SecretValue `json:"data"`
}

type List struct {
	// these fields are needed to validate the manifest
	Kind string `json:"kind"`

	// these are the only data that will be translated
	Items []Secret `json:"items"`
}
