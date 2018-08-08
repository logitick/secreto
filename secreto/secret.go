package secreto

type Secret struct {
	Kind    string            `yaml:"kind"`
	Version string            `yaml:"apiVersion"`
	Type    string            `yaml:"type"`
	Data    map[string]string `yaml:"data"`
}
