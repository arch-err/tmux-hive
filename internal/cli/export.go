package cli

import (
	"fmt"
	"os"

	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/arch-err/tmux-hive/internal/tmux"
	"github.com/spf13/cobra"
)

var exportOutput string

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export current tmux session to a hive configuration",
	Long: `Export the current tmux session to a hive configuration file.

Captures the current session structure including:
- Windows and their layouts
- Panes and their working directories
- Running commands in each pane
- Session options
- Environment variables

Must be run from within a tmux session.`,
	RunE: runExport,
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&exportOutput, "output", "o", "", "output file (default: stdout)")
}

func runExport(cmd *cobra.Command, args []string) error {
	// Check if we're in a tmux session
	sessionName, err := tmux.GetCurrentSession()
	if err != nil {
		logger.Error("Not in a tmux session")
		logger.Info("Run this command from within a tmux session")
		return err
	}

	logger.Infof("Exporting session '%s'", sessionName)

	// Export the session
	cfg, err := tmux.Export()
	if err != nil {
		logger.Error("Failed to export session")
		return err
	}

	// Marshal to YAML
	data, err := config.Marshal(cfg)
	if err != nil {
		logger.Error("Failed to marshal config")
		return err
	}

	// Output to file or stdout
	if exportOutput != "" {
		if err := os.WriteFile(exportOutput, data, 0644); err != nil {
			logger.Error("Failed to write output file")
			return err
		}
		logger.Infof("✓ Exported config written to %s", exportOutput)
	} else {
		fmt.Print(string(data))
	}

	return nil
}
