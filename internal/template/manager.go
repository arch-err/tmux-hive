package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arch-err/tmux-hive/internal/config"
)

// GetTemplateDir returns the directory where templates are stored
// Uses XDG_DATA_HOME/hive/templates or ~/.local/share/hive/templates
func GetTemplateDir() (string, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}

	return filepath.Join(dataHome, "hive", "templates"), nil
}

// EnsureTemplateDir creates the template directory if it doesn't exist
func EnsureTemplateDir() error {
	dir, err := GetTemplateDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create template directory: %w", err)
	}

	return nil
}

// ListTemplates returns a list of available templates
// Includes both built-in templates and user templates
func ListTemplates() ([]string, error) {
	templates := []string{}

	// Add built-in templates
	for name := range builtinTemplates {
		templates = append(templates, name)
	}

	// Add user templates from XDG_DATA_HOME
	dir, err := GetTemplateDir()
	if err != nil {
		return templates, nil // Return built-in templates only
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return templates, nil // Directory doesn't exist yet, return built-in templates
		}
		return nil, fmt.Errorf("failed to read template directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") {
			templateName := strings.TrimSuffix(strings.TrimSuffix(name, ".yaml"), ".yml")
			templates = append(templates, templateName)
		}
	}

	return templates, nil
}

// ReadTemplate reads a template by name
// Checks built-in templates first, then user templates
func ReadTemplate(name string) (*config.Config, error) {
	// Check built-in templates
	if builtinFn, ok := builtinTemplates[name]; ok {
		return builtinFn(), nil
	}

	// Check user templates
	dir, err := GetTemplateDir()
	if err != nil {
		return nil, fmt.Errorf("template not found: %s", name)
	}

	// Try .yaml extension first
	path := filepath.Join(dir, name+".yaml")
	if _, err := os.Stat(path); err == nil {
		return config.Parse(path)
	}

	// Try .yml extension
	path = filepath.Join(dir, name+".yml")
	if _, err := os.Stat(path); err == nil {
		return config.Parse(path)
	}

	return nil, fmt.Errorf("template not found: %s", name)
}

// SaveTemplate saves a config as a template
func SaveTemplate(name string, cfg *config.Config) error {
	if err := EnsureTemplateDir(); err != nil {
		return err
	}

	dir, err := GetTemplateDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, name+".yaml")
	return config.Write(cfg, path)
}

// DeleteTemplate deletes a user template
func DeleteTemplate(name string) error {
	// Don't allow deleting built-in templates
	if _, ok := builtinTemplates[name]; ok {
		return fmt.Errorf("cannot delete built-in template: %s", name)
	}

	dir, err := GetTemplateDir()
	if err != nil {
		return err
	}

	// Try both .yaml and .yml extensions
	path := filepath.Join(dir, name+".yaml")
	if err := os.Remove(path); err == nil {
		return nil
	}

	path = filepath.Join(dir, name+".yml")
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("template not found: %s", name)
	}

	return nil
}
