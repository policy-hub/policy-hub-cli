package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultConfigFilename = "policy-hub"
	envPrefix             = "POLICYHUB"
)

// NewRootCommand returns the root command for the policy-hub CLI
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "policy-hub <command>",
		Short: "Search through Rego policies",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return initializeConfig(cmd)
		},
		SilenceUsage: true,
	}

	return rootCmd
}

// initializeConfig initializes viper config and binds viper to the cobra flags.
func initializeConfig(cmd *cobra.Command) error {

	viper.SetConfigName(defaultConfigFilename)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// If there is no ConfigFile, ignore the error
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			return fmt.Errorf("read config: %w", err)
		}
	}

	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd)

	return nil
}

// bindFlags binds each cobra flag to the associated viper configuration
// Since environment variables can't contain dashes, dashes are translated to underscores
func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
