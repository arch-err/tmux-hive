package config

import (
	"fmt"
	"strings"
)

// ValidationError represents a configuration validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("validation errors:\n")
	for _, err := range e {
		sb.WriteString(fmt.Sprintf("  - %s\n", err.Error()))
	}
	return sb.String()
}

// Validate checks if a Config is valid
func Validate(cfg *Config) error {
	var errors ValidationErrors

	// Validate session
	if cfg.Session.Name == "" {
		errors = append(errors, ValidationError{
			Field:   "session.name",
			Message: "session name is required",
		})
	}

	// Validate windows
	if len(cfg.Windows) == 0 {
		errors = append(errors, ValidationError{
			Field:   "windows",
			Message: "at least one window is required",
		})
	}

	for i, window := range cfg.Windows {
		// Validate window name
		if window.Name == "" {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("windows[%d].name", i),
				Message: "window name is required",
			})
		}

		// Validate layout if specified
		if window.Layout != "" && !isValidLayout(window.Layout) {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("windows[%d].layout", i),
				Message: fmt.Sprintf("invalid layout '%s', must be one of: %s", window.Layout, strings.Join(ValidLayouts, ", ")),
			})
		}

		// Validate panes
		if len(window.Panes) == 0 {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("windows[%d].panes", i),
				Message: "at least one pane is required",
			})
		}

		for j, pane := range window.Panes {
			// Validate split if specified
			if pane.Split != "" && !isValidSplit(pane.Split) {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("windows[%d].panes[%d].split", i, j),
					Message: fmt.Sprintf("invalid split '%s', must be one of: %s", pane.Split, strings.Join(ValidSplits, ", ")),
				})
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func isValidLayout(layout string) bool {
	for _, valid := range ValidLayouts {
		if layout == valid {
			return true
		}
	}
	return false
}

func isValidSplit(split string) bool {
	for _, valid := range ValidSplits {
		if split == valid {
			return true
		}
	}
	return false
}
