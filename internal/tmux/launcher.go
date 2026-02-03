package tmux

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/arch-err/tmux-hive/internal/config"
)

// Launch creates a tmux session from a configuration
func Launch(cfg *config.Config) error {
	// Check if session already exists
	if SessionExists(cfg.Session.Name) {
		return fmt.Errorf("session '%s' already exists. Kill it first or use a different name", cfg.Session.Name)
	}

	// Create the session
	if err := CreateSession(cfg.Session.Name, cfg.Session.BaseDir, cfg.Options); err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	// Set environment variables
	if err := SetEnvVars(cfg.Session.Name, cfg.Env); err != nil {
		return fmt.Errorf("failed to set environment variables: %w", err)
	}

	// Get the base directory for resolving relative paths
	baseDir := cfg.Session.BaseDir
	if baseDir == "" {
		baseDir = "."
	}

	// Track first window index for later selection
	var firstWindowIndex string

	// Create windows and panes
	for i, window := range cfg.Windows {
		windowDir := resolveDir(baseDir, window.Dir)

		var windowIndex string
		if i == 0 {
			// First window is automatically created, get its index
			windows, err := ListWindows(cfg.Session.Name)
			if err != nil {
				return fmt.Errorf("failed to list windows: %w", err)
			}
			if len(windows) == 0 {
				return fmt.Errorf("no windows found in session")
			}
			windowIndex = windows[0].Index
			firstWindowIndex = windowIndex

			// Rename the first window
			if err := renameWindow(cfg.Session.Name, windowIndex, window.Name); err != nil {
				return fmt.Errorf("failed to rename first window: %w", err)
			}
		} else {
			// Create additional windows
			if _, err := CreateWindow(cfg.Session.Name, window.Name, windowDir, window.Layout); err != nil {
				return fmt.Errorf("failed to create window '%s': %w", window.Name, err)
			}
			windowIndex = fmt.Sprintf("%d", i)
		}

		// Create panes
		if len(window.Panes) == 0 {
			continue
		}

		// First pane is already created with the window
		firstPane := window.Panes[0]
		firstPaneDir := resolveDir(windowDir, firstPane.Dir)

		// Get the first pane ID
		panes, err := ListPanes(cfg.Session.Name, windowIndex)
		if err != nil {
			return fmt.Errorf("failed to list panes: %w", err)
		}
		if len(panes) == 0 {
			return fmt.Errorf("no panes found in window")
		}
		firstPaneID := panes[0].ID

		// Change directory if needed
		if firstPaneDir != "" && firstPaneDir != windowDir {
			if err := SendCommand(firstPaneID, fmt.Sprintf("cd %q", firstPaneDir)); err != nil {
				return fmt.Errorf("failed to change directory in first pane: %w", err)
			}
		}

		// Send command to first pane
		if firstPane.Cmd != "" {
			if err := SendCommand(firstPaneID, firstPane.Cmd); err != nil {
				return fmt.Errorf("failed to send command to first pane: %w", err)
			}
		}

		// Create additional panes
		for j := 1; j < len(window.Panes); j++ {
			pane := window.Panes[j]
			paneDir := resolveDir(windowDir, pane.Dir)

			// Create the pane
			paneID, err := CreatePane(cfg.Session.Name, windowIndex, paneDir, pane.Split)
			if err != nil {
				return fmt.Errorf("failed to create pane %d in window '%s': %w", j, window.Name, err)
			}

			// Send command if specified
			if pane.Cmd != "" {
				if err := SendCommand(paneID, pane.Cmd); err != nil {
					return fmt.Errorf("failed to send command to pane: %w", err)
				}
			}
		}

		// Set window layout after all panes are created
		if window.Layout != "" {
			if err := SetWindowLayout(cfg.Session.Name, windowIndex, window.Layout); err != nil {
				return fmt.Errorf("failed to set window layout: %w", err)
			}
		}
	}

	// Select first window
	if len(cfg.Windows) > 0 && firstWindowIndex != "" {
		if err := selectWindow(cfg.Session.Name, firstWindowIndex); err != nil {
			return fmt.Errorf("failed to select first window: %w", err)
		}
	}

	return nil
}

// resolveDir resolves a directory path relative to a base directory
func resolveDir(baseDir, dir string) string {
	if dir == "" {
		return baseDir
	}
	if filepath.IsAbs(dir) {
		return dir
	}
	return filepath.Join(baseDir, dir)
}

// renameWindow renames a window
func renameWindow(sessionName, windowIndex, newName string) error {
	target := fmt.Sprintf("%s:%s", sessionName, windowIndex)
	cmd := exec.Command("tmux", "rename-window", "-t", target, newName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to rename window: %w", err)
	}
	return nil
}

// selectWindow selects a window
func selectWindow(sessionName, windowIndex string) error {
	target := fmt.Sprintf("%s:%s", sessionName, windowIndex)
	cmd := exec.Command("tmux", "select-window", "-t", target)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to select window: %w", err)
	}
	return nil
}
