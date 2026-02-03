package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// SessionExists checks if a tmux session with the given name exists
func SessionExists(name string) bool {
	cmd := exec.Command("tmux", "has-session", "-t", name)
	return cmd.Run() == nil
}

// CreateSession creates a new tmux session
func CreateSession(name, baseDir string, options map[string]interface{}) error {
	args := []string{"new-session", "-d", "-s", name}

	// Set the starting directory if provided
	if baseDir != "" {
		args = append(args, "-c", baseDir)
	}

	cmd := exec.Command("tmux", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	// Apply session options
	for key, value := range options {
		if err := SetSessionOption(name, key, value); err != nil {
			return fmt.Errorf("failed to set option %s: %w", key, err)
		}
	}

	return nil
}

// SetSessionOption sets a tmux session option
func SetSessionOption(sessionName, key string, value interface{}) error {
	var valueStr string
	switch v := value.(type) {
	case string:
		valueStr = v
	case bool:
		if v {
			valueStr = "on"
		} else {
			valueStr = "off"
		}
	case int:
		valueStr = fmt.Sprintf("%d", v)
	default:
		valueStr = fmt.Sprintf("%v", v)
	}

	cmd := exec.Command("tmux", "set-option", "-t", sessionName, key, valueStr)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set option: %w", err)
	}

	return nil
}

// SetEnvVars sets environment variables for a tmux session
func SetEnvVars(sessionName string, env map[string]string) error {
	for key, value := range env {
		cmd := exec.Command("tmux", "set-environment", "-t", sessionName, key, value)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to set environment variable %s: %w", key, err)
		}
	}
	return nil
}

// KillSession kills a tmux session
func KillSession(name string) error {
	cmd := exec.Command("tmux", "kill-session", "-t", name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to kill session: %w", err)
	}
	return nil
}

// GetCurrentSession returns the name of the current tmux session
// Returns empty string if not in a tmux session
func GetCurrentSession() (string, error) {
	cmd := exec.Command("tmux", "display-message", "-p", "#{session_name}")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("not in a tmux session or tmux is not running")
	}

	return strings.TrimSpace(string(output)), nil
}
