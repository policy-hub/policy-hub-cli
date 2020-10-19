package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/policy-hub/policy-hub-cli/pkg/metaconfig"
	"github.com/policy-hub/policy-hub-cli/pkg/search"
	"github.com/spf13/cobra"
)

func newSearchCommand() *cobra.Command {
	c := &searchConfig{}
	cmd := &cobra.Command{
		Use:   "search <flags> <query>",
		Short: "Search through Rego policies",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.run(args[0])
		},
	}

	cmd.Flags().StringVarP(&c.Repository, "repository", "r", "metadata/registries.yml", "Location of the repository to search.")

	return cmd
}

type searchConfig struct {
	Repository string
}

func (c *searchConfig) run(query string) error {
	metadata, err := c.loadMetadata()
	if err != nil {
		return fmt.Errorf("load metadata: %w", err)
	}

	e, err := search.Load()
	if err != nil {
		return fmt.Errorf("load engine: %w", err)
	}

	if err := e.Index(metadata); err != nil {
		return fmt.Errorf("index metadata: %w", err)
	}

	res, err := e.Query(query)
	if err != nil {
		return fmt.Errorf("query engine: %w", err)
	}

	fmt.Println(res)
	return nil
}

func (c *searchConfig) loadMetadata() ([]metaconfig.Metadata, error) {
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
