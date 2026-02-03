# Getting Started

This guide will help you install Hive and create your first tmux session.

## Installation

### Download Binary (Recommended)

Download the latest release for your platform from the [releases page](https://github.com/arch-err/tmux-hive/releases):

```bash
# Linux
curl -L https://github.com/arch-err/tmux-hive/releases/latest/download/tmux-hive_linux_amd64.tar.gz | tar xz
sudo mv hive /usr/local/bin/

# macOS
curl -L https://github.com/arch-err/tmux-hive/releases/latest/download/tmux-hive_darwin_amd64.tar.gz | tar xz
sudo mv hive /usr/local/bin/
```

### Using Go

If you have Go installed:

```bash
go install github.com/arch-err/tmux-hive/cmd/hive@latest
```

### From Source

```bash
git clone https://github.com/arch-err/tmux-hive.git
cd tmux-hive
go build -o hive ./cmd/hive
sudo mv hive /usr/local/bin/
```

## Verify Installation

```bash
hive version
```

## Prerequisites

- **tmux** must be installed and available in your PATH
- For best experience, use tmux 3.0 or later

## Your First Session

### 1. Create a Configuration File

Generate a configuration from a template:

```bash
hive generate -t basic -o .hive.yaml
```

This creates a `.hive.yaml` file in the current directory with a basic configuration.

### 2. Customize the Configuration

Edit the file in your favorite editor:

```bash
hive config
```

Or open it manually:

```bash
vim .hive.yaml
```

Modify the session name and add your commands:

```yaml
session:
  name: my-first-session
  base_dir: .

windows:
  - name: main
    panes:
      - echo "Hello from Hive!"
```

### 3. Validate the Configuration

Check that your configuration is valid:

```bash
hive validate
```

### 4. Launch the Session

Create the tmux session:

```bash
hive launch
```

### 5. Attach to the Session

```bash
tmux attach -t my-first-session
```

## Next Steps

- Explore [built-in templates](commands.md#generate)
- Learn about [configuration options](configuration.md)
- Check out [real-world examples](examples.md)

## Common Workflows

### Development Environment

1. Generate a dev template:
```bash
hive generate -t dev -o .hive.yaml
```

2. Customize for your project
3. Launch with `hive launch`

### CTF Challenge Development

1. Generate a CTF template:
```bash
hive generate -t ctf -o .hive.yaml
```

2. Adjust directories and commands
3. Launch with `hive launch`

### Export Existing Session

Already have a tmux session you like?

1. Attach to your session:
```bash
tmux attach -t my-session
```

2. Export it:
```bash
hive export -o my-session.yaml
```

3. Use it as a template for future sessions

## Troubleshooting

### "tmux: command not found"

Install tmux first:

```bash
# Ubuntu/Debian
sudo apt-get install tmux

# macOS
brew install tmux

# Fedora
sudo dnf install tmux

# Arch Linux
sudo pacman -S tmux
```

### "Session already exists"

If you get an error that the session already exists, kill it first:

```bash
tmux kill-session -t session-name
```

Or use a different session name in your config.

### "No config file found"

Make sure you're in the directory containing `.hive.yaml` or `hive.yaml`, or specify the config path:

```bash
hive launch -c /path/to/config.yaml
```
