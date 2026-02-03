# Welcome to Hive 🐝

Hive is a modern tmux session manager that uses YAML configuration files to define and launch tmux sessions with windows and panes.

## Why Hive?

- **Simple**: YAML configuration is easy to read and write
- **Fast**: Written in Go, single binary with no dependencies
- **Beautiful**: Modern CLI with interactive prompts powered by Charmbracelet
- **Portable**: Works on Linux, macOS, and Windows (with WSL)
- **Flexible**: Built-in templates and export from existing sessions

## Key Features

### YAML Configuration
Define your entire tmux session layout in a simple YAML file:

```yaml
session:
  name: my-project
  base_dir: ~/projects/my-project

windows:
  - name: editor
    panes:
      - nvim .

  - name: server
    panes:
      - npm run dev
```

### Built-in Templates
Quick start with pre-configured templates for common workflows:

- **basic** - Single window, single pane
- **dev** - Editor + terminal + logs
- **ctf** - CTF challenge development setup
- **web** - Frontend + backend + database

### Export & Import
Export your existing tmux sessions to configuration files:

```bash
tmux attach -t my-session
hive export -o my-session.yaml
```

### Interactive CLI
Beautiful interactive prompts for template selection and configuration:

```bash
hive generate
# Interactive template picker appears
```

## Quick Links

- [Getting Started](getting-started.md) - Installation and first steps
- [Configuration](configuration.md) - Full configuration reference
- [Commands](commands.md) - CLI command reference
- [Examples](examples.md) - Real-world configuration examples

## Community

- [GitHub Repository](https://github.com/arch-err/tmux-hive)
- [Issue Tracker](https://github.com/arch-err/tmux-hive/issues)
- [Discussions](https://github.com/arch-err/tmux-hive/discussions)

## License

Hive is released under the MIT License. See [LICENSE](https://github.com/arch-err/tmux-hive/blob/main/LICENSE) for details.
