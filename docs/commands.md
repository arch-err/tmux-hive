# Commands Reference

Complete reference for all Hive CLI commands.

## Global Flags

These flags are available for all commands:

- `-c, --config <path>` - Specify config file path
- `-v, --verbose` - Enable verbose logging
- `-h, --help` - Show help

## hive generate

Generate a hive configuration file from a template.

### Usage

```bash
hive generate [flags]
```

### Flags

- `-t, --template <name>` - Template to use
- `-o, --output <file>` - Output file (default: stdout)

### Examples

Interactive template selection:
```bash
hive generate -o .hive.yaml
```

Use specific template:
```bash
hive generate -t dev -o .hive.yaml
```

Output to stdout:
```bash
hive generate -t basic
```

### Built-in Templates

- **basic** - Single window, single pane (minimal setup)
- **dev** - Development environment (editor + terminal + logs)
- **ctf** - CTF challenge development (editor + docker + recon + notes)
- **web** - Web development (frontend + backend + database)
- **blank/minimal** - Blank template to start from scratch

### User Templates

User templates are stored in `~/.local/share/hive/templates/` (or `$XDG_DATA_HOME/hive/templates`).

To create a user template:
1. Create your config file
2. Save it with a `.yaml` extension in the templates directory
3. Use it with `hive generate -t <name>`

## hive launch

Launch a tmux session from a hive configuration.

### Usage

```bash
hive launch [flags]
```

### Flags

None specific. Uses global flags.

### Examples

Launch from `.hive.yaml` or `hive.yaml`:
```bash
hive launch
```

Launch from specific file:
```bash
hive launch -c my-config.yaml
```

### Notes

- Session must not already exist
- Config file must be valid (run `hive validate` first if unsure)
- Creates session in detached mode
- Use `tmux attach -t <session-name>` to attach

## hive export

Export current tmux session to a hive configuration.

### Usage

```bash
hive export [flags]
```

### Flags

- `-o, --output <file>` - Output file (default: stdout)

### Examples

Export to stdout:
```bash
hive export
```

Export to file:
```bash
hive export -o my-session.yaml
```

Export current session from within tmux:
```bash
tmux attach -t my-session
hive export -o my-session.yaml
```

### What Gets Exported

- Session name
- Window names and layouts
- Pane structure and working directories
- Running commands in each pane
- Session options
- Environment variables

### Notes

- Must be run from within a tmux session
- Commands are captured as currently running (may differ from how they were started)
- Layout names may not be preserved exactly (exports layout structure)

## hive validate

Validate a hive configuration file.

### Usage

```bash
hive validate [flags]
```

### Flags

None specific. Uses global flags.

### Examples

Validate `.hive.yaml`:
```bash
hive validate
```

Validate specific file:
```bash
hive validate -c my-config.yaml
```

### What Gets Validated

- YAML syntax
- Required fields (session name, windows, panes)
- Valid option values (layouts, split directions)
- Field types and formats

### Output

Success:
```
INFO Validating .hive.yaml
INFO ✓ Configuration is valid
```

Errors:
```
INFO Validating .hive.yaml
ERROR Validation failed
validation errors:
  - session.name: session name is required
  - windows[0].panes: at least one pane is required
```

## hive config

Edit the hive configuration file in `$EDITOR`.

### Usage

```bash
hive config [flags]
```

### Flags

None specific. Uses global flags.

### Examples

Edit `.hive.yaml`:
```bash
hive config
```

Edit specific file:
```bash
hive config -c my-config.yaml
```

### Notes

- Opens config in `$EDITOR` (falls back to `vi`)
- Validates config after editing
- Reports validation errors if any
- If no config exists, prompts to generate one

## hive version

Show version information.

### Usage

```bash
hive version
```

### Example Output

```
hive version v0.1.0
Commit: abc123
Built: 2024-01-01T00:00:00Z
Go: go1.23.0
OS/Arch: linux/amd64
```

## Common Workflows

### Create New Project Session

```bash
# 1. Navigate to project directory
cd ~/projects/my-project

# 2. Generate config from template
hive generate -t dev -o .hive.yaml

# 3. Customize config
hive config

# 4. Validate config
hive validate

# 5. Launch session
hive launch

# 6. Attach to session
tmux attach -t <session-name>
```

### Export and Reuse Session

```bash
# 1. Create your ideal tmux setup manually
tmux new -s my-session
# ... create windows and panes ...

# 2. Export it
hive export -o my-session.yaml

# 3. Kill the test session
tmux kill-session -t my-session

# 4. Launch from config
hive launch -c my-session.yaml
```

### Share Configuration

```bash
# 1. Create and test your config
hive generate -t web -o .hive.yaml
hive config
hive launch

# 2. Commit to git
git add .hive.yaml
git commit -m "Add hive config"
git push

# 3. Team members can use it
git pull
hive launch
```

### Template Creation

```bash
# 1. Create a great config
hive generate -t basic -o my-template.yaml
hive config -c my-template.yaml

# 2. Save as user template
mkdir -p ~/.local/share/hive/templates
cp my-template.yaml ~/.local/share/hive/templates/

# 3. Use it for new projects
cd ~/new-project
hive generate -t my-template -o .hive.yaml
```
