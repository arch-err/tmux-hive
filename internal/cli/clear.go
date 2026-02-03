package cli

import (
	"fmt"

	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/arch-err/tmux-hive/internal/tmux"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Kill the tmux session defined in the config",
	Long: `Kill the tmux session defined in the hive configuration file.

Asks for confirmation before killing the session.`,
	RunE: runClear,
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

func runClear(cmd *cobra.Command, args []string) error {
	// Discover config file
	configPath, err := config.Discover(cfgFile)
	if err != nil {
		logger.Error("No config file found")
		logger.Info("Run 'hive generate' to create a new config file")
		return err
	}

	// Parse config to get session name
	cfg, err := config.Parse(configPath)
	if err != nil {
		logger.Error("Failed to parse config")
		return err
	}

	// Check if session exists
	if !tmux.SessionExists(cfg.Session.Name) {
		logger.Infof("Session '%s' does not exist", cfg.Session.Name)
		return nil
	}

	// Ask for confirmation
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("Kill session '%s'?", cfg.Session.Name)).
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
	return nil
}
