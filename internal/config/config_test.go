package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestPaneConfigUnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		expected PaneConfig
	}{
		{
			name: "string format",
			yaml: `"echo hello"`,
			expected: PaneConfig{
				Cmd: "echo hello",
			},
		},
		{
			name: "empty string",
			yaml: `""`,
			expected: PaneConfig{
				Cmd: "",
			},
		},
		{
			name: "object format with cmd",
			yaml: `
cmd: npm run dev
dir: ./frontend
split: horizontal`,
			expected: PaneConfig{
				Cmd:   "npm run dev",
				Dir:   "./frontend",
				Split: "horizontal",
			},
		},
		{
			name: "object format with only cmd",
			yaml: `
cmd: ls -la`,
			expected: PaneConfig{
				Cmd: "ls -la",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pane PaneConfig
			err := yaml.Unmarshal([]byte(tt.yaml), &pane)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if pane.Cmd != tt.expected.Cmd {
				t.Errorf("Cmd: got %q, want %q", pane.Cmd, tt.expected.Cmd)
			}
			if pane.Dir != tt.expected.Dir {
				t.Errorf("Dir: got %q, want %q", pane.Dir, tt.expected.Dir)
			}
			if pane.Split != tt.expected.Split {
				t.Errorf("Split: got %q, want %q", pane.Split, tt.expected.Split)
			}
		})
	}
}

func TestValidLayouts(t *testing.T) {
	expected := []string{
		"even-horizontal",
		"even-vertical",
		"main-horizontal",
		"main-vertical",
		"tiled",
	}

	if len(ValidLayouts) != len(expected) {
		t.Errorf("ValidLayouts length: got %d, want %d", len(ValidLayouts), len(expected))
	}

	for i, layout := range expected {
		if ValidLayouts[i] != layout {
			t.Errorf("ValidLayouts[%d]: got %q, want %q", i, ValidLayouts[i], layout)
		}
	}
}

func TestValidSplits(t *testing.T) {
	expected := []string{
		"horizontal",
		"vertical",
	}

	if len(ValidSplits) != len(expected) {
		t.Errorf("ValidSplits length: got %d, want %d", len(ValidSplits), len(expected))
	}

	for i, split := range expected {
		if ValidSplits[i] != split {
			t.Errorf("ValidSplits[%d]: got %q, want %q", i, ValidSplits[i], split)
		}
	}
}
