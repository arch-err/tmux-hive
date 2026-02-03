package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscover(t *testing.T) {
	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	defer os.Chdir(originalDir)

	t.Run("explicit config path", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "custom.yaml")

		if err := os.WriteFile(configPath, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		result, err := Discover(configPath)
		if err != nil {
			t.Fatalf("Discover() error = %v", err)
		}

		if result != configPath {
			t.Errorf("Discover() = %q, want %q", result, configPath)
		}
	})

	t.Run("explicit config path not found", func(t *testing.T) {
		_, err := Discover("/nonexistent/config.yaml")
		if err == nil {
			t.Error("Discover() should return error for nonexistent file")
		}
	})

	t.Run("find .hive.yaml", func(t *testing.T) {
		tmpDir := t.TempDir()
		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		dotHivePath := ".hive.yaml"
		if err := os.WriteFile(dotHivePath, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create .hive.yaml: %v", err)
		}

		result, err := Discover("")
		if err != nil {
			t.Fatalf("Discover() error = %v", err)
		}

		if result != dotHivePath {
			t.Errorf("Discover() = %q, want %q", result, dotHivePath)
		}
	})

	t.Run("find hive.yaml", func(t *testing.T) {
		tmpDir := t.TempDir()
		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		hivePath := "hive.yaml"
		if err := os.WriteFile(hivePath, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create hive.yaml: %v", err)
		}

		result, err := Discover("")
		if err != nil {
			t.Fatalf("Discover() error = %v", err)
		}

		if result != hivePath {
			t.Errorf("Discover() = %q, want %q", result, hivePath)
		}
	})

	t.Run("prefer .hive.yaml over hive.yaml", func(t *testing.T) {
		tmpDir := t.TempDir()
		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		dotHivePath := ".hive.yaml"
		hivePath := "hive.yaml"

		if err := os.WriteFile(dotHivePath, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create .hive.yaml: %v", err)
		}
		if err := os.WriteFile(hivePath, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create hive.yaml: %v", err)
		}

		result, err := Discover("")
		if err != nil {
			t.Fatalf("Discover() error = %v", err)
		}

		if result != dotHivePath {
			t.Errorf("Discover() = %q, want %q (should prefer .hive.yaml)", result, dotHivePath)
		}
	})

	t.Run("no config found", func(t *testing.T) {
		tmpDir := t.TempDir()
		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		_, err := Discover("")
		if err == nil {
			t.Error("Discover() should return error when no config found")
		}
	})
}

func TestDiscoverAbs(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	dotHivePath := ".hive.yaml"
	if err := os.WriteFile(dotHivePath, []byte("test"), 0644); err != nil {
		t.Fatalf("failed to create .hive.yaml: %v", err)
	}

	result, err := DiscoverAbs("")
	if err != nil {
		t.Fatalf("DiscoverAbs() error = %v", err)
	}

	if !filepath.IsAbs(result) {
		t.Errorf("DiscoverAbs() = %q, should be absolute path", result)
	}

	expectedAbs := filepath.Join(tmpDir, dotHivePath)
	if result != expectedAbs {
		t.Errorf("DiscoverAbs() = %q, want %q", result, expectedAbs)
	}
}
