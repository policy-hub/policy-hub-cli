package metaconfig_test

import (
	"io/ioutil"
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

		t.Run("should support metadata for multiple packages in a single file", func(t *testing.T) {
			assert.Greater(t, len(config), 1)
		})

		t.Run("should support a description field", func(t *testing.T) {
			assert.NotEmpty(t, config[0].Description)
			assert.Equal(t, config[0].Description, "this is a description field")
		})

		t.Run("should support a version field", func(t *testing.T) {
			assert.NotEmpty(t, config[0].Version)
			assert.Equal(t, config[0].Version, "v1")
		})

		t.Run("should support a labels field", func(t *testing.T) {
			assert.NotEmpty(t, config[0].Labels)
			assert.Contains(t, config[0].Labels, "aws")
			assert.Contains(t, config[0].Labels, "security")
			assert.NotContains(t, config[0].Labels, "random-nonesense-that-doesnt-exist")
		})
	})
}
