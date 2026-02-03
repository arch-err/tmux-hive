package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseBytes(t *testing.T) {
	tests := []struct {
		name    string
		yaml    string
		wantErr bool
	}{
		{
			name: "valid minimal config",
			yaml: `
session:
  name: test
windows:
  - name: main
    panes:
      - echo hello`,
			wantErr: false,
		},
		{
			name: "valid complex config",
			yaml: `
session:
  name: test
  base_dir: /tmp
windows:
  - name: editor
    dir: ./src
    layout: main-vertical
    panes:
      - cmd: nvim .
      - cmd: ""
        split: vertical
options:
  mouse: on
  base-index: 1
env:
  DEBUG: "true"`,
			wantErr: false,
		},
		{
			name: "invalid yaml",
			yaml: `
session:
  name: test
windows:
  - name: main
    panes:
      - [invalid yaml`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := ParseBytes([]byte(tt.yaml))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && cfg == nil {
				t.Error("ParseBytes() returned nil config")
			}
		})
	}
}

func TestParse(t *testing.T) {
	// Create temporary test file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test.yaml")

	configContent := `
session:
  name: test-session
  base_dir: /tmp
windows:
  - name: main
    panes:
      - echo hello
options:
  mouse: on
env:
  TEST: "value"`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	cfg, err := Parse(configPath)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if cfg.Session.Name != "test-session" {
		t.Errorf("Session.Name: got %q, want %q", cfg.Session.Name, "test-session")
	}

	if cfg.Session.BaseDir != "/tmp" {
		t.Errorf("Session.BaseDir: got %q, want %q", cfg.Session.BaseDir, "/tmp")
	}

	if len(cfg.Windows) != 1 {
		t.Errorf("Windows length: got %d, want %d", len(cfg.Windows), 1)
	}

	if cfg.Windows[0].Name != "main" {
		t.Errorf("Windows[0].Name: got %q, want %q", cfg.Windows[0].Name, "main")
	}
}

func TestParseNonExistentFile(t *testing.T) {
	_, err := Parse("/nonexistent/file.yaml")
	if err == nil {
		t.Error("Parse() should return error for nonexistent file")
	}
}

func TestMarshal(t *testing.T) {
	cfg := &Config{
		Session: SessionConfig{
			Name:    "test",
			BaseDir: "/tmp",
		},
		Windows: []WindowConfig{
			{
				Name: "main",
				Panes: []PaneConfig{
					{Cmd: "echo hello"},
				},
			},
		},
		Options: map[string]interface{}{
			"mouse": "on",
		},
		Env: map[string]string{
			"TEST": "value",
		},
	}

	data, err := Marshal(cfg)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("Marshal() returned empty data")
	}

	// Verify we can parse it back
	parsed, err := ParseBytes(data)
	if err != nil {
		t.Fatalf("ParseBytes() after Marshal() error = %v", err)
	}

	if parsed.Session.Name != cfg.Session.Name {
		t.Errorf("Round-trip Session.Name: got %q, want %q", parsed.Session.Name, cfg.Session.Name)
	}
}

func TestWrite(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test.yaml")

	cfg := &Config{
		Session: SessionConfig{
			Name: "test",
		},
		Windows: []WindowConfig{
			{
				Name: "main",
				Panes: []PaneConfig{
					{Cmd: "echo hello"},
				},
			},
		},
	}

	err := Write(cfg, configPath)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Verify file exists and can be read
	parsed, err := Parse(configPath)
	if err != nil {
		t.Fatalf("Parse() after Write() error = %v", err)
	}

	if parsed.Session.Name != cfg.Session.Name {
		t.Errorf("Session.Name after Write(): got %q, want %q", parsed.Session.Name, cfg.Session.Name)
	}
}
