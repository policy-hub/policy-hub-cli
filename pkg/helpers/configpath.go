package helpers

import (
	"os"
	"path/filepath"
)

const PolicyIndexDirName = ".policy-hub-index"
const PolicyDirName = ".policy-hub"
const PolicyRegistryFilename = "registries.list"
const RegistryMetadataDir = "metadata"
const DefaultRegistriesFilename = "registries.yml"
const RegistryListDefault = "https://raw.githubusercontent.com/policy-hub/policy-hub-cli/main/metadata/" + DefaultRegistriesFilename

//DefaultMetaDataFile  Returns the path to the default registries yaml for the current system
func DefaultMetaDataFile() string {
	return filepath.ToSlash(filepath.Join(
		MetadataConfigPath(), DefaultRegistriesFilename))
}

//MetadataConfigPath  Returns the path to the metadata directory for the current system
func MetadataConfigPath() string {
	directory := filepath.ToSlash(filepath.Join(ConfigPath(), RegistryMetadataDir))
	return directory
}

//ConfigPath  Returns the path to the config directory for the current system
func ConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	directory := filepath.ToSlash(filepath.Join(homeDir, PolicyDirName))
	return directory
}

//IndexPath  Returns the path to the search index directory for the current system
func IndexPath() string {
	homeDir, _ := os.UserHomeDir()
	directory := filepath.ToSlash(filepath.Join(homeDir, PolicyIndexDirName))
	return directory
}
