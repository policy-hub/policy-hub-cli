package metaconfig

import "gopkg.in/yaml.v2"

type MetadataConfig struct {
	Description string
	Version     string
	Labels      []string
}

func Load(config []byte) ([]MetadataConfig, error) {
	c := make([]MetadataConfig, 0)
	err := yaml.Unmarshal(config, &c)
	return c, err
}
