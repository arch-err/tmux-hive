package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// CreatePane creates a new pane by splitting an existing pane
// Returns the pane ID of the newly created pane
func CreatePane(sessionName, windowIndex, dir, split string) (string, error) {
	target := fmt.Sprintf("%s:%s", sessionName, windowIndex)
	args := []string{"split-window", "-t", target, "-P", "-F", "#{pane_id}"}

	// Set split direction
	if split == "horizontal" {
		args = append(args, "-h")
	} else if split == "vertical" {
		args = append(args, "-v")
	} else {
		// Default to vertical split
		args = append(args, "-v")
	}

	// Set starting directory
	if dir != "" {
		args = append(args, "-c", dir)
	}

	cmd := exec.Command("tmux", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to create pane: %w", err)
	}

	paneID := strings.TrimSpace(string(output))
	return paneID, nil
}

// SendCommand sends a command to a pane
func SendCommand(paneID, command string) error {
	if command == "" {
		return nil
	}

	// Send the command followed by Enter
	cmd := exec.Command("tmux", "send-keys", "-t", paneID, command, "Enter")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to send command to pane: %w", err)
	}

	return nil
}

// ListPanes returns a list of panes in a window
func ListPanes(sessionName, windowIndex string) ([]PaneInfo, error) {
	target := fmt.Sprintf("%s:%s", sessionName, windowIndex)
	cmd := exec.Command("tmux", "list-panes", "-t", target, "-F", "#{pane_id}:#{pane_current_path}:#{pane_current_command}")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list panes: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	panes := make([]PaneInfo, 0, len(lines))

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 3)
		if len(parts) == 3 {
			panes = append(panes, PaneInfo{
				ID:      parts[0],
				Dir:     parts[1],
				Command: parts[2],
			})
		}
	}

	return panes, nil
}

// PaneInfo contains information about a tmux pane
type PaneInfo struct {
	ID      string
	Dir     string
	Command string
}
