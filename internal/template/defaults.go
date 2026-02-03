package template

import "github.com/arch-err/tmux-hive/internal/config"

// builtinTemplates maps template names to their generator functions
var builtinTemplates = map[string]func() *config.Config{
	"basic":  basicTemplate,
	"dev":    devTemplate,
	"ctf":    ctfTemplate,
	"web":    webTemplate,
	"blank":  blankTemplate,
	"minimal": blankTemplate,
}

// basicTemplate returns a basic single-window, single-pane configuration
func basicTemplate() *config.Config {
	return &config.Config{
		Session: config.SessionConfig{
			Name:    "basic",
			BaseDir: ".",
		},
		Windows: []config.WindowConfig{
			{
				Name: "main",
				Panes: []config.PaneConfig{
					{Cmd: ""},
				},
			},
		},
		Options: map[string]interface{}{
			"mouse":         "on",
			"base-index":    1,
			"history-limit": 50000,
		},
	}
}

// devTemplate returns a development environment configuration
func devTemplate() *config.Config {
	return &config.Config{
		Session: config.SessionConfig{
			Name:    "dev",
			BaseDir: ".",
		},
		Windows: []config.WindowConfig{
			{
				Name:   "editor",
				Layout: "main-vertical",
				Panes: []config.PaneConfig{
					{Cmd: "nvim ."},
					{Cmd: "", Split: "vertical"},
				},
			},
			{
				Name:   "terminal",
				Layout: "even-vertical",
				Panes: []config.PaneConfig{
					{Cmd: ""},
					{Cmd: "", Split: "vertical"},
				},
			},
			{
				Name: "logs",
				Panes: []config.PaneConfig{
					{Cmd: ""},
				},
			},
		},
		Options: map[string]interface{}{
			"mouse":         "on",
			"base-index":    1,
			"history-limit": 50000,
		},
	}
}

// ctfTemplate returns a CTF challenge development configuration
func ctfTemplate() *config.Config {
	return &config.Config{
		Session: config.SessionConfig{
			Name:    "ctf-dev",
			BaseDir: ".",
		},
		Windows: []config.WindowConfig{
			{
				Name:   "editor",
				Layout: "main-vertical",
				Panes: []config.PaneConfig{
					{Cmd: "nvim ."},
					{Cmd: "", Split: "vertical"},
				},
			},
			{
				Name: "docker",
				Dir:  "./docker",
				Panes: []config.PaneConfig{
					{Cmd: "docker-compose up"},
					{Cmd: "docker-compose logs -f", Split: "horizontal"},
				},
			},
			{
				Name:   "recon",
				Layout: "tiled",
				Panes: []config.PaneConfig{
					{Cmd: ""},
					{Cmd: "", Split: "vertical"},
					{Cmd: "", Split: "horizontal"},
					{Cmd: "", Split: "vertical"},
				},
			},
			{
				Name: "notes",
				Panes: []config.PaneConfig{
					{Cmd: ""},
				},
			},
		},
		Options: map[string]interface{}{
			"mouse":         "on",
			"base-index":    1,
			"history-limit": 50000,
		},
		Env: map[string]string{
			"CHALLENGE_ID": "challenge-001",
			"DEBUG":        "true",
		},
	}
}

// webTemplate returns a web development configuration
func webTemplate() *config.Config {
	return &config.Config{
		Session: config.SessionConfig{
			Name:    "web-dev",
			BaseDir: ".",
		},
		Windows: []config.WindowConfig{
			{
				Name:   "editor",
				Layout: "main-vertical",
				Panes: []config.PaneConfig{
					{Cmd: "nvim ."},
					{Cmd: "", Split: "vertical"},
				},
			},
			{
				Name: "frontend",
				Dir:  "./frontend",
				Panes: []config.PaneConfig{
					{Cmd: "npm run dev"},
					{Cmd: "npm run test -- --watch", Split: "vertical"},
				},
			},
			{
				Name: "backend",
				Dir:  "./backend",
				Panes: []config.PaneConfig{
					{Cmd: "npm run dev"},
					{Cmd: "npm run test -- --watch", Split: "vertical"},
				},
			},
			{
				Name: "database",
				Panes: []config.PaneConfig{
					{Cmd: "docker-compose up postgres redis"},
					{Cmd: "", Split: "vertical"},
				},
			},
		},
		Options: map[string]interface{}{
			"mouse":         "on",
			"base-index":    1,
			"history-limit": 50000,
		},
		Env: map[string]string{
			"NODE_ENV": "development",
		},
	}
}

// blankTemplate returns a minimal blank configuration
func blankTemplate() *config.Config {
	return &config.Config{
		Session: config.SessionConfig{
			Name:    "my-session",
			BaseDir: ".",
		},
		Windows: []config.WindowConfig{
			{
				Name: "main",
				Panes: []config.PaneConfig{
					{Cmd: ""},
				},
			},
		},
		Options: map[string]interface{}{
			"mouse":         "on",
			"base-index":    1,
			"history-limit": 50000,
		},
	}
}
