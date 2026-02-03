package config

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid minimal config",
			config: &Config{
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
			},
			wantErr: false,
		},
		{
			name: "missing session name",
			config: &Config{
				Session: SessionConfig{},
				Windows: []WindowConfig{
					{
						Name: "main",
						Panes: []PaneConfig{
							{Cmd: "echo hello"},
						},
					},
				},
			},
			wantErr: true,
			errMsg:  "session.name",
		},
		{
			name: "no windows",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{},
			},
			wantErr: true,
			errMsg:  "at least one window is required",
		},
		{
			name: "window without name",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{
					{
						Panes: []PaneConfig{
							{Cmd: "echo hello"},
						},
					},
				},
			},
			wantErr: true,
			errMsg:  "window name is required",
		},
		{
			name: "window without panes",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{
					{
						Name:  "main",
						Panes: []PaneConfig{},
					},
				},
			},
			wantErr: true,
			errMsg:  "at least one pane is required",
		},
		{
			name: "invalid layout",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{
					{
						Name:   "main",
						Layout: "invalid-layout",
						Panes: []PaneConfig{
							{Cmd: "echo hello"},
						},
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid layout",
		},
		{
			name: "valid layout",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{
					{
						Name:   "main",
						Layout: "main-vertical",
						Panes: []PaneConfig{
							{Cmd: "echo hello"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid split",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{
					{
						Name: "main",
						Panes: []PaneConfig{
							{Cmd: "echo hello"},
							{Cmd: "echo world", Split: "invalid-split"},
						},
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid split",
		},
		{
			name: "valid split",
			config: &Config{
				Session: SessionConfig{
					Name: "test",
				},
				Windows: []WindowConfig{
					{
						Name: "main",
						Panes: []PaneConfig{
							{Cmd: "echo hello"},
							{Cmd: "echo world", Split: "horizontal"},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != "" {
				if err == nil || !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Validate() error should contain %q, got %v", tt.errMsg, err)
				}
			}
		})
	}
}

func TestIsValidLayout(t *testing.T) {
	tests := []struct {
		layout string
		valid  bool
	}{
		{"even-horizontal", true},
		{"even-vertical", true},
		{"main-horizontal", true},
		{"main-vertical", true},
		{"tiled", true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.layout, func(t *testing.T) {
			result := isValidLayout(tt.layout)
			if result != tt.valid {
				t.Errorf("isValidLayout(%q) = %v, want %v", tt.layout, result, tt.valid)
			}
		})
	}
}

func TestIsValidSplit(t *testing.T) {
	tests := []struct {
		split string
		valid bool
	}{
		{"horizontal", true},
		{"vertical", true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.split, func(t *testing.T) {
			result := isValidSplit(tt.split)
			if result != tt.valid {
				t.Errorf("isValidSplit(%q) = %v, want %v", tt.split, result, tt.valid)
			}
		})
	}
}
