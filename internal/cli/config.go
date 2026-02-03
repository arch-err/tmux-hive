package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit the hive configuration file",
	Long: `Edit the hive configuration file in $EDITOR.

After editing, the config will be validated automatically.
If no config file exists, you'll be prompted to generate one.`,
	RunE: runConfig,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func runConfig(cmd *cobra.Command, args []string) error {
	// Discover config file
	configPath, err := config.Discover(cfgFile)
	if err != nil {
		logger.Error("No config file found")
		logger.Info("Run 'hive generate' to create a new config file")
		return err
	}

	// Get editor from environment
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi" // Fallback to vi
	}

	logger.Infof("Opening %s in %s", configPath, editor)

	// Open editor
	editorCmd := exec.Command(editor, configPath)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		logger.Error("Failed to run editor")
		return err
	}

	// Validate config after editing
	logger.Info("Validating configuration...")
	cfg, err := config.Parse(configPath)
	if err != nil {
		logger.Error("Failed to parse config")
		return err
	}

	if err := config.Validate(cfg); err != nil {
		logger.Error("Validation failed")
		fmt.Println(err.Error())
		return err
	}

	logger.Info("✓ Configuration is valid")
	return nil
}
