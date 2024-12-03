package model

type Component struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		Annotations map[string]string `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		Type      string `yaml:"type"`
		Owner     string `yaml:"owner"`
		Lifecycle string `yaml:"lifecycle"`
	} `yaml:"spec"`
}
