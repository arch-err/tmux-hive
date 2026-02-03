package cli

import (
	"os"
	"os/exec"

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

	// Check if we're already in a tmux session
	inTmux := os.Getenv("TMUX") != ""

	var switchCmd *exec.Cmd
	if inTmux {
		// Switch to the new session instead of attaching
		switchCmd = exec.Command("tmux", "switch-client", "-t", cfg.Session.Name)
	} else {
		// Attach to the session
		switchCmd = exec.Command("tmux", "attach", "-t", cfg.Session.Name)
	}

	switchCmd.Stdin = os.Stdin
	switchCmd.Stdout = os.Stdout
	switchCmd.Stderr = os.Stderr

	if err := switchCmd.Run(); err != nil {
		logger.Warnf("Failed to attach/switch to session (session created successfully): %v", err)
		if inTmux {
			logger.Infof("Switch manually with: tmux switch-client -t %s", cfg.Session.Name)
		} else {
			logger.Infof("Attach manually with: tmux attach -t %s", cfg.Session.Name)
		}
	}

	return nil
}
