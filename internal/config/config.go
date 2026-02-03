package config

// Config represents the complete hive configuration
type Config struct {
	Session SessionConfig          `yaml:"session"`
	Windows []WindowConfig         `yaml:"windows"`
	Options map[string]interface{} `yaml:"options,omitempty"`
	Env     map[string]string      `yaml:"env,omitempty"`
}

// SessionConfig represents session-level configuration
type SessionConfig struct {
	Name    string `yaml:"name"`
	BaseDir string `yaml:"base_dir,omitempty"`
}

// WindowConfig represents a tmux window configuration
type WindowConfig struct {
	Name   string       `yaml:"name"`
	Dir    string       `yaml:"dir,omitempty"`
	Layout string       `yaml:"layout,omitempty"`
	Panes  []PaneConfig `yaml:"panes"`
}

// PaneConfig represents a tmux pane configuration
// Can be specified as a string (command only) or as a struct with additional options
type PaneConfig struct {
	Cmd   string `yaml:"cmd,omitempty"`
	Dir   string `yaml:"dir,omitempty"`
	Split string `yaml:"split,omitempty"` // "horizontal" or "vertical"
}

// UnmarshalYAML implements custom unmarshaling for PaneConfig
// Allows panes to be specified as either a string or a struct
func (p *PaneConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Try to unmarshal as a string first
	var cmd string
	if err := unmarshal(&cmd); err == nil {
		p.Cmd = cmd
		return nil
	}

	// If that fails, try to unmarshal as a struct
	type paneAlias PaneConfig
	var pane paneAlias
	if err := unmarshal(&pane); err != nil {
		return err
	}

	*p = PaneConfig(pane)
	return nil
}

// ValidLayouts are the supported tmux layout strings
var ValidLayouts = []string{
	"even-horizontal",
	"even-vertical",
	"main-horizontal",
	"main-vertical",
	"tiled",
}

// ValidSplits are the supported pane split directions
var ValidSplits = []string{
	"horizontal",
	"vertical",
}
