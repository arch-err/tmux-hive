# Configuration Reference

This page provides a complete reference for Hive configuration files.

## File Location

Hive looks for configuration files in the following order:

1. Path specified with `-c/--config` flag
2. `.hive.yaml` in the current directory
3. `hive.yaml` in the current directory

## Configuration Structure

```yaml
session:
  # Session configuration

windows:
  # List of windows

options:
  # Tmux options

env:
  # Environment variables
```

## Session Configuration

The `session` section defines session-level settings.

### `session.name` (required)

The name of the tmux session.

```yaml
session:
  name: my-project
```

### `session.base_dir` (optional)

The base directory for the session. All relative paths in window and pane directories will be resolved relative to this path.

```yaml
session:
  name: my-project
  base_dir: ~/projects/my-project
```

## Windows Configuration

The `windows` section is a list of window definitions.

### `windows[].name` (required)

The name of the window.

```yaml
windows:
  - name: editor
```

### `windows[].dir` (optional)

The working directory for this window. Overrides `session.base_dir`.

```yaml
windows:
  - name: editor
    dir: ~/projects/my-project/src
```

### `windows[].layout` (optional)

The tmux layout for this window. Valid values:

- `even-horizontal` - Panes are spread out evenly horizontally
- `even-vertical` - Panes are spread out evenly vertically
- `main-horizontal` - One large pane on top, others below
- `main-vertical` - One large pane on left, others on right
- `tiled` - Panes are arranged in a grid

```yaml
windows:
  - name: editor
    layout: main-vertical
```

### `windows[].panes` (required)

A list of pane definitions for this window. At least one pane is required.

```yaml
windows:
  - name: editor
    panes:
      - nvim .
      - # empty pane
```

## Pane Configuration

Panes can be defined in two ways:

### Simple String Format

For panes that only need a command:

```yaml
panes:
  - echo "Hello"
  - ls -la
  - # empty pane (just opens shell)
```

### Object Format

For panes that need additional configuration:

```yaml
panes:
  - cmd: echo "Hello"
    dir: ./subdirectory
    split: horizontal
```

### `panes[].cmd` (optional)

The command to run in this pane. If omitted, just opens a shell.

```yaml
panes:
  - cmd: npm run dev
```

### `panes[].dir` (optional)

The working directory for this pane. Overrides `windows[].dir` and `session.base_dir`.

```yaml
panes:
  - cmd: npm run dev
    dir: ./frontend
```

### `panes[].split` (optional)

How to split the window to create this pane. Only applies to panes after the first one.

Valid values:
- `horizontal` - Split horizontally (side-by-side)
- `vertical` - Split vertically (top-bottom)

```yaml
panes:
  - cmd: nvim .
  - cmd: npm test
    split: horizontal  # Creates pane to the right
  - cmd: npm run build
    split: vertical    # Creates pane below
```

## Options Configuration

The `options` section allows you to set tmux options for the session.

### Common Options

```yaml
options:
  mouse: on                    # Enable mouse support
  base-index: 1                # Start window numbering at 1
  history-limit: 50000         # Scrollback buffer size
  status-position: top         # Status bar position (top/bottom)
  default-terminal: screen-256color  # Terminal type
```

### All Tmux Options

You can use any valid tmux option. See `man tmux` for a complete list.

## Environment Variables

The `env` section defines environment variables for the session.

```yaml
env:
  NODE_ENV: development
  DEBUG: "true"
  API_KEY: your-api-key
  PORT: "3000"
```

!!! warning
    Environment variables are set at the session level and inherited by all windows and panes.

## Complete Example

```yaml
session:
  name: web-project
  base_dir: ~/projects/web-app

windows:
  - name: editor
    layout: main-vertical
    panes:
      - cmd: nvim .
      - cmd: ""
        split: vertical

  - name: frontend
    dir: ./frontend
    panes:
      - cmd: npm run dev
      - cmd: npm run test -- --watch
        split: horizontal

  - name: backend
    dir: ./backend
    panes:
      - cmd: npm run dev
        dir: .
      - cmd: npm run test -- --watch
        split: horizontal

  - name: database
    panes:
      - cmd: docker-compose up postgres
      - cmd: docker-compose logs -f postgres
        split: vertical

  - name: logs
    layout: tiled
    panes:
      - cmd: tail -f logs/app.log
      - cmd: tail -f logs/error.log
        split: horizontal
      - cmd: tail -f logs/access.log
        split: vertical
      - cmd: htop
        split: horizontal

options:
  mouse: on
  base-index: 1
  history-limit: 50000
  status-position: top
  default-terminal: screen-256color

env:
  NODE_ENV: development
  DEBUG: "true"
  DATABASE_URL: postgresql://localhost/myapp
  API_PORT: "3000"
  FRONTEND_PORT: "8080"
```

## Best Practices

### Use Base Directory

Set `session.base_dir` and use relative paths for windows and panes:

```yaml
session:
  base_dir: ~/projects/my-app

windows:
  - name: frontend
    dir: ./frontend  # Relative to base_dir
```

### Empty Panes

Leave `cmd` empty for panes where you want to run commands manually:

```yaml
panes:
  - nvim .
  - # Empty pane for manual commands
```

### Named Sessions

Use descriptive session names to easily identify them:

```yaml
session:
  name: myapp-dev  # Not just "dev"
```

### Layouts

Choose layouts based on your workflow:

- `main-vertical` - Good for editor + terminal
- `tiled` - Good for monitoring multiple services
- `even-horizontal` - Good for side-by-side comparisons

### Comments

Use YAML comments to document your configuration:

```yaml
windows:
  - name: services
    panes:
      - docker-compose up  # Start all services
      - docker-compose logs -f  # Follow logs
```
