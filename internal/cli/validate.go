package cli

import (
	"fmt"

	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a hive configuration file",
	Long: `Validate a hive configuration file for syntax and semantic errors.

Checks for:
- Valid YAML syntax
- Required fields (session name, windows, etc.)
- Valid option values (layouts, split directions)`,
	RunE: runValidate,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func runValidate(cmd *cobra.Command, args []string) error {
	// Discover config file
	configPath, err := config.Discover(cfgFile)
	if err != nil {
		logger.Error("No config file found")
		logger.Info("Run 'hive generate' to create a new config file")
		return err
	}

	logger.Infof("Validating %s", configPath)

	// Parse config
	cfg, err := config.Parse(configPath)
	if err != nil {
		logger.Error("Failed to parse config")
		return err
	}

	// Validate config
	if err := config.Validate(cfg); err != nil {
		logger.Error("Validation failed")
		fmt.Println(err.Error())
		return err
	}

	logger.Info("✓ Configuration is valid")
	return nil
}
