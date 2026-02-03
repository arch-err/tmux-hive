# Roadmap

This document outlines the current features and planned future enhancements for Hive.

## ✅ Completed (v0.1.0)

### Core Features
- [x] YAML configuration parsing
- [x] Session creation and management
- [x] Window creation with layouts
- [x] Pane creation and splitting
- [x] Command execution in panes
- [x] Environment variable support
- [x] Session options configuration

### CLI Commands
- [x] `hive generate` - Generate config from templates
- [x] `hive launch` - Launch session from config
- [x] `hive export` - Export current session to config
- [x] `hive validate` - Validate config file
- [x] `hive config` - Edit config in $EDITOR
- [x] `hive version` - Show version info

### Template System
- [x] Built-in templates (basic, dev, ctf, web, blank)
- [x] User templates in XDG_DATA_HOME/hive/templates
- [x] Interactive template selection
- [x] Template save and load

### Export Capabilities
- [x] Capture window structure
- [x] Capture pane layout
- [x] Capture running commands
- [x] Capture working directories
- [x] Capture session options
- [x] Capture environment variables

### Developer Experience
- [x] Beautiful CLI with Charmbracelet tools
- [x] Interactive prompts
- [x] Helpful error messages
- [x] Configuration validation
- [x] GitHub Actions CI/CD
- [x] GoReleaser for multi-platform builds

## 🚧 In Progress

### Documentation
- [ ] Complete MkDocs documentation
- [ ] Usage examples and tutorials
- [ ] Configuration reference
- [ ] Best practices guide

### Testing
- [ ] Unit tests for core packages
- [ ] Integration tests
- [ ] End-to-end tests

## 📋 Planned Features

### Priority 1: Essential Features

#### Sync Command
- [ ] `hive sync` - Synchronize current session with config
- [ ] Detect differences between config and running session
- [ ] Smart reconciliation (add/remove windows/panes)
- [ ] Confirm before destructive operations
- [ ] Preserve running processes when possible

#### Wait Conditions & Dependencies
- [ ] Pane dependencies (wait for command to succeed before starting next)
- [ ] Health checks (wait for port to be available, file to exist, etc.)
- [ ] Timeout configuration
- [ ] Retry logic
- [ ] Example: Don't start frontend until backend is healthy

```yaml
windows:
  - name: backend
    panes:
      - cmd: npm run dev
        health_check:
          port: 3000
          timeout: 30s

  - name: frontend
    depends_on:
      - backend
    panes:
      - cmd: npm run dev
```

#### Hooks System
- [ ] Pre/post hooks for sessions
- [ ] Pre/post hooks for windows
- [ ] Pre/post hooks for panes
- [ ] Environment setup hooks
- [ ] Cleanup hooks

```yaml
hooks:
  before_session:
    - docker-compose up -d
  after_session:
    - echo "Session ready!"
  before_window:
    backend:
      - npm install
  after_window:
    backend:
      - echo "Backend window ready"
```

### Priority 2: Advanced Configuration

#### Variable Substitution
- [ ] Environment variable substitution (`${VAR}`)
- [ ] Built-in variables (`${PROJECT_NAME}`, `${USER}`, etc.)
- [ ] Config-defined variables
- [ ] Conditional substitution
- [ ] Default values

```yaml
variables:
  project_name: my-app
  port: 3000

session:
  name: ${project_name}

windows:
  - name: server
    panes:
      - PORT=${port} npm run dev
```

#### Config Includes
- [ ] Include other config files
- [ ] Compose multiple configs
- [ ] Override mechanisms
- [ ] Template inheritance

```yaml
include:
  - base-dev.yaml
  - docker-services.yaml

session:
  name: my-project  # Override from base
```

#### Profiles
- [ ] Multiple configuration profiles
- [ ] Switch between dev/debug/demo modes
- [ ] Profile-specific window sets
- [ ] Profile-specific commands

```yaml
profiles:
  dev:
    windows:
      - name: editor
      - name: terminal

  demo:
    windows:
      - name: presentation
      - name: demo-app
```

### Priority 3: Enhanced Workflow

#### Watch Mode
- [ ] `hive watch` - Auto-reload on config changes
- [ ] File watching with debouncing
- [ ] Selective reloading (changed windows only)
- [ ] Notification on reload

#### Session Management
- [ ] `hive list` - List all hive-managed sessions
- [ ] `hive attach` - Attach to a session
- [ ] `hive down` - Kill session defined in config
- [ ] `hive restart` - Restart a session
- [ ] Session tagging and filtering

#### Shell Completions
- [ ] Bash completion
- [ ] Zsh completion
- [ ] Fish completion
- [ ] PowerShell completion

### Priority 4: Integration & Compatibility

#### Import from Other Tools
- [ ] Import tmuxinator configs
- [ ] Import tmuxp configs
- [ ] Migration guide
- [ ] Format conversion tool

#### Remote Sessions
- [ ] SSH integration
- [ ] Create sessions on remote hosts
- [ ] Remote config management
- [ ] Connection pooling

```yaml
session:
  name: remote-dev
  host: dev-server.example.com
  user: developer
```

#### Integration with Other Tools
- [ ] Git integration (branch detection, repo info)
- [ ] Docker integration (container status, logs)
- [ ] Kubernetes integration (pod logs, exec)
- [ ] Systemd integration (service status)

### Priority 5: Advanced Features

#### Pane Management
- [ ] Pane focus configuration
- [ ] Pane zoom settings
- [ ] Pane size percentages
- [ ] Pane titles/labels
- [ ] Pane borders customization

#### Layout Management
- [ ] Custom layout strings
- [ ] Layout templates
- [ ] Responsive layouts (adapt to terminal size)
- [ ] Layout presets library

#### Logging & Monitoring
- [ ] Pane output logging to files
- [ ] Aggregate logs from all panes
- [ ] Log rotation
- [ ] Log filtering and searching

#### Interactive Mode
- [ ] `hive interactive` - Build config interactively
- [ ] Step-by-step wizard
- [ ] Preview before creation
- [ ] Save as template

### Priority 6: Ecosystem

#### Plugin System
- [ ] Plugin architecture
- [ ] Plugin discovery
- [ ] Plugin manager
- [ ] Community plugins

#### GUI / TUI
- [ ] Terminal UI for session management
- [ ] Visual config editor
- [ ] Session monitoring dashboard
- [ ] Real-time pane output viewer

#### API / Library
- [ ] Go library for programmatic use
- [ ] REST API for remote management
- [ ] Webhook support
- [ ] Event system

## 🤔 Ideas Under Consideration

- Session persistence across reboots
- Cloud config storage and sync
- Session sharing and collaboration
- AI-assisted config generation
- Performance profiling and optimization
- Config validation in CI/CD
- Session templates marketplace
- Integration with terminal multiplexers other than tmux

## 📝 Notes

- Features are prioritized based on user needs and implementation complexity
- Timelines are estimates and may change based on contributions and feedback
- Community input is welcome on feature priorities

## Contributing

Want to help implement these features? Check out our [contributing guide](CONTRIBUTING.md) and pick an issue labeled with `help wanted` or `good first issue`.
