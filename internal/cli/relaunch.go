package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/arch-err/tmux-hive/internal/tmux"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var relaunchCmd = &cobra.Command{
	Use:   "relaunch",
	Short: "Kill and relaunch the tmux session",
	Long: `Kill the existing tmux session (if it exists) and relaunch it from the config.

Combines 'hive clear' and 'hive launch' into a single command.
Asks for confirmation before killing the existing session.`,
	RunE: runRelaunch,
}

func init() {
	rootCmd.AddCommand(relaunchCmd)
}

func runRelaunch(cmd *cobra.Command, args []string) error {
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

	// Check if session exists
	if tmux.SessionExists(cfg.Session.Name) {
		// Ask for confirmation to kill
		var confirm bool
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title(fmt.Sprintf("Kill existing session '%s' and relaunch?", cfg.Session.Name)).
					Description("This will terminate the session and all processes running in it.").
					Value(&confirm),
			),
		)

		if err := form.Run(); err != nil {
			return fmt.Errorf("confirmation cancelled")
		}

		if !confirm {
			logger.Info("Cancelled")
			return nil
		}

		// Kill the session
		logger.Infof("Killing session '%s'", cfg.Session.Name)
		if err := tmux.KillSession(cfg.Session.Name); err != nil {
			logger.Error("Failed to kill session")
			return err
		}

		logger.Infof("✓ Session '%s' killed", cfg.Session.Name)
	}

	// Launch the session (same as launch command)
	logger.Infof("Launching session '%s'", cfg.Session.Name)

	if err := tmux.Launch(cfg); err != nil {
		logger.Error("Failed to launch session")
		return err
	}

	logger.Infof("✓ Session '%s' launched successfully", cfg.Session.Name)

	// Attach to the session
	attachCmd := exec.Command("tmux", "attach", "-t", cfg.Session.Name)
	attachCmd.Stdin = os.Stdin
	attachCmd.Stdout = os.Stdout
	attachCmd.Stderr = os.Stderr

	if err := attachCmd.Run(); err != nil {
		logger.Warnf("Failed to attach to session (session created successfully): %v", err)
		logger.Infof("Attach manually with: tmux attach -t %s", cfg.Session.Name)
	}

	return nil
}
