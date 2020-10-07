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

		t.Run("should support a description field", func(t *testing.T) {
			assert.NotEmpty(t, config[0].Description)
			assert.Equal(t, config[0].Description, "this is a description field")
		})

		t.Run("should support a version field", func(t *testing.T) {
			assert.NotEmpty(t, config[0].Version)
			assert.Equal(t, config[0].Version, "v1")
		})
	})
}
