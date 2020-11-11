package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/open-policy-agent/conftest/downloader"
	"github.com/policy-hub/policy-hub-cli/internal/commands"
	"github.com/policy-hub/policy-hub-cli/pkg/helpers"
)

func init() {
	checkCreateDir()
	checkCreateRegistriesList()
	checkCreateRegistriesMetadata()
}

func main() {
	if err := commands.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func checkCreateRegistriesMetadata() {
	metadataPath := helpers.MetadataConfigPath()
	_, err := os.Stat(metadataPath)
	if os.IsNotExist(err) {

		registriesListPath := filepath.ToSlash(
			filepath.Join(helpers.ConfigPath(), helpers.PolicyRegistryFilename))
		content, err := ioutil.ReadFile(registriesListPath)
		if err != nil {
			log.Panic("could not read registries list file: %w", err)
		}

		lines := strings.Split(string(content), "\n")
		err = downloader.Download(context.Background(), metadataPath, lines)
		if err != nil {
			log.Panic("could not fetch metadata file: %w", err)
		}
	}
}

func checkCreateRegistriesList() {
	configPath := helpers.ConfigPath()
	registriesListPath := filepath.ToSlash(
		filepath.Join(configPath, helpers.PolicyRegistryFilename))

	_, err := os.Stat(registriesListPath)
	if os.IsNotExist(err) {
		file, err := os.Create(registriesListPath)
		defer file.Close()
		if err != nil {
			log.Panic("could not create file: %w", err)
		}

		_, err = io.WriteString(file, helpers.RegistryListDefault)
		if err != nil {
			log.Panic("could not write to file: %w", err)
		}

		file.Sync()
	}
}

func checkCreateDir() {
	configPath := helpers.ConfigPath()
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		os.MkdirAll(configPath, os.ModePerm)
	}
}
