# Hive 🐝

[![Build](https://github.com/arch-err/tmux-hive/actions/workflows/build.yml/badge.svg)](https://github.com/arch-err/tmux-hive/actions/workflows/build.yml)
[![Test](https://github.com/arch-err/tmux-hive/actions/workflows/test.yml/badge.svg)](https://github.com/arch-err/tmux-hive/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arch-err/tmux-hive)](https://goreportcard.com/report/github.com/arch-err/tmux-hive)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern tmux session manager with YAML configuration. Similar to tmuxinator and tmuxp, but written in Go with a beautiful CLI experience powered by [Charmbracelet](https://github.com/charmbracelet).

## Features

- 📝 **YAML Configuration**: Define your tmux sessions with simple, readable YAML
- 🎨 **Beautiful CLI**: Modern terminal UI with interactive prompts
- 📦 **Built-in Templates**: Quick start with pre-configured templates (basic, dev, ctf, web)
- 🔄 **Export & Import**: Export existing tmux sessions to configs
- ✅ **Validation**: Catch errors before launching
- 🚀 **Fast & Portable**: Single binary, no dependencies

## Installation

### Download Binary

Download the latest release for your platform from the [releases page](https://github.com/arch-err/tmux-hive/releases).

```bash
# Linux/macOS
curl -L https://github.com/arch-err/tmux-hive/releases/latest/download/tmux-hive_<version>_<os>_<arch>.tar.gz | tar xz
sudo mv hive /usr/local/bin/

# Or using go install
go install github.com/arch-err/tmux-hive/cmd/hive@latest
```

### From Source

```bash
git clone https://github.com/arch-err/tmux-hive.git
cd tmux-hive
go build -o hive ./cmd/hive
sudo mv hive /usr/local/bin/
```

## Quick Start

### 1. Generate a Configuration

```bash
# Interactive template selection
hive generate -o .hive.yaml

# Or specify a template directly
hive generate -t dev -o .hive.yaml
```

### 2. Edit Your Configuration

```bash
hive config
```

### 3. Launch Your Session

```bash
hive launch
```

## Example Configuration

```yaml
session:
  name: my-project
  base_dir: ~/projects/my-project

windows:
  - name: editor
    layout: main-vertical
    panes:
      - nvim .
      - # empty pane for terminal

  - name: servers
    dir: .
    panes:
      - npm run dev
      - docker-compose up

  - name: logs
    layout: tiled
    panes:
      - tail -f logs/app.log
      - tail -f logs/error.log

options:
  mouse: on
  base-index: 1
  history-limit: 50000

env:
  NODE_ENV: development
  DEBUG: "true"
```

## Commands

- `hive generate` - Generate a config from a template
- `hive launch` - Launch a tmux session from config
- `hive export` - Export current tmux session to config
- `hive validate` - Validate a config file
- `hive config` - Edit the config file
- `hive version` - Show version information

## Documentation

For full documentation, visit [https://arch-err.github.io/tmux-hive/](https://arch-err.github.io/tmux-hive/)

## Roadmap

See [ROADMAP.md](ROADMAP.md) for planned features.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- UI powered by [Charmbracelet](https://github.com/charmbracelet)
- Inspired by [tmuxinator](https://github.com/tmuxinator/tmuxinator) and [tmuxp](https://github.com/tmux-python/tmuxp)
