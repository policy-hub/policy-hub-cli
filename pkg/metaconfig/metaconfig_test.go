package metaconfig_test

import (
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/policy-hub/policy-hub-cli/pkg/metaconfig"
	"github.com/stretchr/testify/assert"
)

func TestMetaConfig(t *testing.T) {
	t.Run("metaconfig/v1", func(t *testing.T) {
		f, err := ioutil.ReadFile("testdata/metadata_v1.yml")
		assert.NoError(t, err, "Error reading fixture")
		config, err := metaconfig.Load(f)
		assert.NoError(t, err, "Error loading config")
		assert.NotEmpty(t, config)
		randomIndex := rand.Intn(len(config) - 1)

		t.Run("should support metadata for multiple packages in a single file", func(t *testing.T) {
			assert.Greater(t, len(config), 1)
		})

		t.Run("should support a homepage field", func(t *testing.T) {
			assert.NotEmpty(t, config[randomIndex].Homepage)
			assert.Equal(t, config[randomIndex].Homepage, "www.policy.io")
		})

		t.Run("should support a path field", func(t *testing.T) {
			assert.NotEmpty(t, config[randomIndex].Path)
			assert.Equal(t, config[randomIndex].Path, "github.com/cool/policies")
		})

		t.Run("should support a description field", func(t *testing.T) {
			assert.NotEmpty(t, config[randomIndex].Description)
			assert.Equal(t, config[randomIndex].Description, "this is a description field")
		})

		t.Run("should support a version field", func(t *testing.T) {
			assert.NotEmpty(t, config[randomIndex].Version)
			assert.Equal(t, config[randomIndex].Version, "v1")
		})

		t.Run("should support a labels field", func(t *testing.T) {
			assert.NotEmpty(t, config[randomIndex].Labels)
			assert.Contains(t, config[randomIndex].Labels, "aws")
			assert.Contains(t, config[randomIndex].Labels, "security")
			assert.NotContains(t, config[randomIndex].Labels, "random-nonesense-that-doesnt-exist")
		})

		t.Run("should support a maintainers field", func(t *testing.T) {
			assert.NotEmpty(t, config[randomIndex].Maintainers)
			assert.Contains(t, config[randomIndex].Maintainers, "person1")
			assert.NotContains(t, config[randomIndex].Maintainers, "random-nonesense-that-doesnt-exist")
		})
	})
}
