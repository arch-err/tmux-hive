package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Discover finds a hive configuration file
// Looks for:
// 1. Path specified in configPath (if not empty)
// 2. .hive.yaml in current directory
// 3. hive.yaml in current directory
func Discover(configPath string) (string, error) {
	// If a config path is explicitly provided, use it
	if configPath != "" {
		if _, err := os.Stat(configPath); err != nil {
			return "", fmt.Errorf("config file not found: %s", configPath)
		}
		return configPath, nil
	}

	// Look for .hive.yaml
	dotHivePath := ".hive.yaml"
	if _, err := os.Stat(dotHivePath); err == nil {
		return dotHivePath, nil
	}

	// Look for hive.yaml
	hivePath := "hive.yaml"
	if _, err := os.Stat(hivePath); err == nil {
		return hivePath, nil
	}

	return "", fmt.Errorf("no hive config file found (.hive.yaml or hive.yaml)")
}

// DiscoverAbs is like Discover but returns an absolute path
func DiscoverAbs(configPath string) (string, error) {
	path, err := Discover(configPath)
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	return absPath, nil
}
