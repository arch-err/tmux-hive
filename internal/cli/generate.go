package cli

import (
	"fmt"
	"os"

	"github.com/arch-err/tmux-hive/internal/config"
	"github.com/arch-err/tmux-hive/internal/template"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var (
	generateTemplate string
	generateOutput   string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a hive configuration file from a template",
	Long: `Generate a hive configuration file from a template.

Interactively select from built-in templates or custom templates stored in
XDG_DATA_HOME/hive/templates (~/.local/share/hive/templates).`,
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&generateTemplate, "template", "t", "", "template to use")
	generateCmd.Flags().StringVarP(&generateOutput, "output", "o", "", "output file (default: stdout)")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	var selectedTemplate string

	// If template is specified via flag, use it
	if generateTemplate != "" {
		selectedTemplate = generateTemplate
	} else {
		// Interactive template selection
		templates, err := template.ListTemplates()
		if err != nil {
			return fmt.Errorf("failed to list templates: %w", err)
		}

		if len(templates) == 0 {
			return fmt.Errorf("no templates available")
		}

		// Create options for huh.Select
		options := make([]huh.Option[string], len(templates))
		for i, t := range templates {
			options[i] = huh.NewOption(t, t)
		}

		var selected string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select a template").
					Options(options...).
					Value(&selected),
			),
		)

		if err := form.Run(); err != nil {
			return fmt.Errorf("template selection cancelled")
		}

		selectedTemplate = selected
	}

	// Read the template
	cfg, err := template.ReadTemplate(selectedTemplate)
	if err != nil {
		return fmt.Errorf("failed to read template: %w", err)
	}

	// Marshal to YAML
	data, err := config.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Output to file or stdout
	if generateOutput != "" {
		if err := os.WriteFile(generateOutput, data, 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		logger.Infof("Generated config written to %s", generateOutput)
	} else {
		fmt.Print(string(data))
	}

	return nil
}
