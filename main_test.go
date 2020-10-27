package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/stretchr/testify/assert"
)

func TestMainCLI(t *testing.T) {
	gomega.RegisterTestingT(t)
	pathToMainCLI, err := gexec.Build("main.go")
	assert.NoError(t, err, "Error creating binary")
	defer gexec.CleanupBuildArtifacts()

	t.Run("we can search registries", func(t *testing.T) {
		t.Skip("the search command seems to hang.")
		command := exec.Command(pathToMainCLI, "search", "k8s", "-r", "./metadata/registries.yml")
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
}
