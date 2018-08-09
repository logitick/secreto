package secreto

type Secret struct {
	// these fields are needed to validate the manifest
	Kind    string `yaml:"kind"`
	Version string `yaml:"apiVersion"`
	Type    string `yaml:"type"`

	// these are the only data that will be translated
	Data map[string]string `yaml:"data"`
}
