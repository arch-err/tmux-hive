package cli

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	logger  *log.Logger
)

var rootCmd = &cobra.Command{
	Use:   "hive",
	Short: "A modern tmux session manager",
	Long:  "\033[1mhive\033[0m",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = log.NewWithOptions(os.Stderr, log.Options{
			ReportCaller:    false,
			ReportTimestamp: false,
		})

		if verbose {
			logger.SetLevel(log.DebugLevel)
		} else {
			logger.SetLevel(log.InfoLevel)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path (.hive.yaml or hive.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
