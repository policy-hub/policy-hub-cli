package commands

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/open-policy-agent/conftest/downloader"
	"github.com/policy-hub/policy-hub-cli/pkg/helpers"
	"github.com/policy-hub/policy-hub-cli/pkg/metaconfig"
	"github.com/spf13/cobra"
)

func newPullCommand() *cobra.Command {
	c := &pullConfig{}
	cmd := &cobra.Command{
		Use:   "pull <flags> <query>",
		Short: "Pull Rego policies to use locally",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.run(args[0])
		},
	}

	cmd.Flags().StringVarP(&c.Repository, "repository", "r", helpers.DefaultMetaDataFile(), "Location of the repository to search.")
	cmd.Flags().StringVarP(&c.PolicyDir, "policy", "p", "", "Folder where the policies will be downloaded too")
	return cmd
}

type pullConfig struct {
	Repository string
	PolicyDir string
}

func (c *pullConfig) run(name string) error {
	metadata, err := c.loadMetadata()
	if err != nil {
		return fmt.Errorf("load metadata: %w", err)
	}

	var urlPath string
	for _, record := range metadata {
		if record.Name == name {
			urlPath = record.Path
		}
	}

	if urlPath == "" {
		return fmt.Errorf("could not find a name match in given repository")
	}

	policyDir := name
	if c.PolicyDir != "" {
		policyDir = c.PolicyDir
	}
	
	err = downloader.Download(context.Background(), policyDir, []string{urlPath})
	return nil
}

func (c *pullConfig) loadMetadata() ([]metaconfig.Metadata, error) {
	f, err := ioutil.ReadFile(c.Repository)
	if err != nil {
		return nil, fmt.Errorf("load repository file: %w", err)
	}

	metadata, err := metaconfig.Load(f)
	if err != nil {
		return nil, fmt.Errorf("load metadata: %w", err)
	}

	return metadata, nil
}
