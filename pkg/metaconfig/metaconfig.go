package metaconfig

import "gopkg.in/yaml.v2"

type Metadata struct {
	Name        string
	Description string
	Version     string
	Labels      []string
	Maintainers []string
	Path        string
	Homepage    string
	ImageURL    string `yaml:"image_url"`
}

func Load(config []byte) ([]Metadata, error) {
	c := make([]Metadata, 0)
	err := yaml.Unmarshal(config, &c)
	return c, err
}
