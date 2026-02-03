package tmux

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/arch-err/tmux-hive/internal/config"
)

// Export captures the current tmux session and converts it to a Config
func Export() (*config.Config, error) {
	// Get current session name
	sessionName, err := GetCurrentSession()
	if err != nil {
		return nil, fmt.Errorf("failed to get current session: %w", err)
	}

	cfg := &config.Config{
		Session: config.SessionConfig{
			Name: sessionName,
		},
		Windows: []config.WindowConfig{},
		Options: make(map[string]interface{}),
		Env:     make(map[string]string),
	}

	// Get session options
	options, err := getSessionOptions(sessionName)
	if err == nil {
		cfg.Options = options
	}

	// Get environment variables
	env, err := getSessionEnv(sessionName)
	if err == nil {
		cfg.Env = env
	}

	// Get windows
	windows, err := ListWindows(sessionName)
	if err != nil {
		return nil, fmt.Errorf("failed to list windows: %w", err)
	}

	for _, window := range windows {
		windowCfg := config.WindowConfig{
			Name:   window.Name,
			Layout: guessLayoutName(window.Layout),
			Panes:  []config.PaneConfig{},
		}

		// Get panes for this window
		panes, err := ListPanes(sessionName, window.Index)
		if err != nil {
			return nil, fmt.Errorf("failed to list panes for window %s: %w", window.Name, err)
		}

		for i, pane := range panes {
			paneCfg := config.PaneConfig{
				Dir: pane.Dir,
			}

			// Try to get the running command (excluding shell)
			if pane.Command != "bash" && pane.Command != "zsh" && pane.Command != "sh" {
				paneCfg.Cmd = pane.Command
			}

			// Set split direction for non-first panes
			if i > 0 {
				// Default to vertical split (can't reliably detect from tmux state)
				paneCfg.Split = "vertical"
			}

			windowCfg.Panes = append(windowCfg.Panes, paneCfg)
		}

		cfg.Windows = append(cfg.Windows, windowCfg)
	}

	return cfg, nil
}

// getSessionOptions retrieves session options
func getSessionOptions(sessionName string) (map[string]interface{}, error) {
	options := make(map[string]interface{})

	// Get commonly used options
	commonOptions := []string{
		"mouse",
		"base-index",
		"history-limit",
	}

	for _, opt := range commonOptions {
		cmd := exec.Command("tmux", "show-options", "-t", sessionName, opt)
		output, err := cmd.Output()
		if err != nil {
			continue // Option might not be set
		}

		// Parse output: "option-name value"
		line := strings.TrimSpace(string(output))
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			value := strings.Join(parts[1:], " ")
			options[opt] = value
		}
	}

	return options, nil
}

// getSessionEnv retrieves session environment variables
func getSessionEnv(sessionName string) (map[string]string, error) {
	env := make(map[string]string)

	cmd := exec.Command("tmux", "show-environment", "-t", sessionName)
	output, err := cmd.Output()
	if err != nil {
		return env, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "-") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	return env, nil
}

// guessLayoutName tries to map a layout string to a named layout
// This is a best-effort approach as tmux stores layouts as complex strings
func guessLayoutName(layoutStr string) string {
	// Tmux layouts are complex strings, we can't reliably determine the name
	// Just return empty for now, user can set it manually
	return ""
}
