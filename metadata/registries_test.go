package metadata_test

import (
	"io/ioutil"
	"testing"

	"github.com/policy-hub/policy-hub-cli/pkg/metaconfig"
	"github.com/stretchr/testify/assert"
)

func TestRegistries(t *testing.T) {
	t.Run("check registries.yml file has valid records", func(t *testing.T) {
		f, err := ioutil.ReadFile("registries.yml")
		assert.NoError(t, err, "Error reading fixture")
		config, err := metaconfig.Load(f)
		assert.NoError(t, err, "Error loading config")
		assert.NotEmpty(t, config)

		t.Run("required fields check", func(t *testing.T) {
			for _, registry := range config {
				assert.NotEmpty(t, registry.Path)
				assert.NotEmpty(t, registry.Name)
				assert.NotEmpty(t, registry.Labels)
				assert.NotEmpty(t, registry.Description)
			}
		})
	})
}
