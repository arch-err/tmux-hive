package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// CreateWindow creates a new window in the specified session
func CreateWindow(sessionName, windowName, dir, layout string) (string, error) {
	args := []string{"new-window", "-t", sessionName, "-n", windowName, "-P", "-F", "#{window_index}"}

	if dir != "" {
		args = append(args, "-c", dir)
	}

	cmd := exec.Command("tmux", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to create window: %w", err)
	}

	windowIndex := strings.TrimSpace(string(output))

	// Set layout if specified
	if layout != "" {
		if err := SetWindowLayout(sessionName, windowIndex, layout); err != nil {
			return "", err
		}
	}

	return windowIndex, nil
}

// SetWindowLayout sets the layout for a window
func SetWindowLayout(sessionName, windowIndex, layout string) error {
	target := fmt.Sprintf("%s:%s", sessionName, windowIndex)
	cmd := exec.Command("tmux", "select-layout", "-t", target, layout)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set layout: %w", err)
	}
	return nil
}

// ListWindows returns a list of windows in a session
func ListWindows(sessionName string) ([]WindowInfo, error) {
	cmd := exec.Command("tmux", "list-windows", "-t", sessionName, "-F", "#{window_index}:#{window_name}:#{window_layout}")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list windows: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	windows := make([]WindowInfo, 0, len(lines))

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 3)
		if len(parts) == 3 {
			windows = append(windows, WindowInfo{
				Index:  parts[0],
				Name:   parts[1],
				Layout: parts[2],
			})
		}
	}

	return windows, nil
}

// WindowInfo contains information about a tmux window
type WindowInfo struct {
	Index  string
	Name   string
	Layout string
}
