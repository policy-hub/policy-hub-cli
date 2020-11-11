package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/policy-hub/policy-hub-cli/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

func TestMainCLI(t *testing.T) {
	gomega.RegisterTestingT(t)
	pathToMainCLI, err := gexec.Build("main.go")
	assert.NoError(t, err, "Error creating binary")
	defer gexec.CleanupBuildArtifacts()
	defer os.RemoveAll(helpers.IndexPath())
	defer os.RemoveAll(helpers.ConfigPath())

	t.Run("we can search registries", func(t *testing.T) {
		command := exec.Command(pathToMainCLI, "search", "k8s")
		session, err := gexec.Start(command, os.Stdout, os.Stdout)
		session.Wait()
		assert.NoError(t, err, "Error running search command")
	})

	t.Run("we can download registries", func(t *testing.T) {
		policyPackageName := "contrib.image_enforcer"
		defer os.RemoveAll(policyPackageName)
		outputSpy := bytes.NewBufferString("")
		command := exec.Command(pathToMainCLI, "pull", policyPackageName)
		session, err := gexec.Start(command, outputSpy, outputSpy)
		session.Wait(10 * time.Second)
		assert.NoError(t, err, "Error running pull command")
		assert.NotContains(t, outputSpy.String(), `Error: unknown command "pull"`)
		_, err = os.Stat(policyPackageName)
		assert.False(t, os.IsNotExist(err), "could not find the directory of policies")
	})

	t.Run("cli should create registry information", func(t *testing.T) {
		configPath := helpers.ConfigPath()
		os.RemoveAll(configPath)
		defer os.RemoveAll(configPath)
		command := exec.Command(pathToMainCLI, "help")
		session, _ := gexec.Start(command, os.Stdout, os.Stdout)
		session.Wait(10 * time.Second)

		t.Run("should create the config directory", func(t *testing.T) {
			_, err := os.Stat(configPath)
			assert.False(t, os.IsNotExist(err), "could not find the directory config directory")
		})

		t.Run("should create the registries list", func(t *testing.T) {
			_, err = os.Stat(filepath.ToSlash(filepath.Join(helpers.ConfigPath(), helpers.PolicyRegistryFilename)))
			assert.False(t, os.IsNotExist(err), "could not find the registries list file")
		})

		t.Run("should create the registries file from the list", func(t *testing.T) {
			_, err = os.Stat(helpers.MetadataConfigPath())
			assert.False(t, os.IsNotExist(err), "could not find the registries metadata files")
		})
	})
}
