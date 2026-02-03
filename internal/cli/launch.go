package cli

import (
	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/arch-err/tmux-hive/internal/tmux"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch a tmux session from a hive configuration",
	Long: `Launch a tmux session from a hive configuration file.

Creates a new tmux session with windows and panes as defined in the config.
If the session already exists, an error will be returned.`,
	RunE: runLaunch,
}

func init() {
	rootCmd.AddCommand(launchCmd)
}

func runLaunch(cmd *cobra.Command, args []string) error {
	// Discover config file
	configPath, err := config.Discover(cfgFile)
	if err != nil {
		logger.Error("No config file found")
		logger.Info("Run 'hive generate' to create a new config file")
		return err
	}

	logger.Infof("Loading config from %s", configPath)

	// Parse config
	cfg, err := config.Parse(configPath)
	if err != nil {
		logger.Error("Failed to parse config")
		return err
	}

	// Validate config
	if err := config.Validate(cfg); err != nil {
		logger.Error("Invalid configuration")
		return err
	}

	// Check if session already exists
	if tmux.SessionExists(cfg.Session.Name) {
		logger.Errorf("Session '%s' already exists", cfg.Session.Name)
		logger.Info("Kill the session first with: tmux kill-session -t %s", cfg.Session.Name)
		return err
	}

	logger.Infof("Launching session '%s'", cfg.Session.Name)

	// Launch the session
	if err := tmux.Launch(cfg); err != nil {
		logger.Error("Failed to launch session")
		return err
	}

	logger.Infof("✓ Session '%s' launched successfully", cfg.Session.Name)
	logger.Infof("Attach to the session with: tmux attach -t %s", cfg.Session.Name)

	return nil
}
