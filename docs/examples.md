# Configuration Examples

Real-world configuration examples for common use cases.

## Basic Development Environment

A simple setup with editor and terminal.

```yaml
session:
  name: basic-dev
  base_dir: ~/projects/my-app

windows:
  - name: editor
    layout: main-vertical
    panes:
      - nvim .
      - # Empty pane for commands

  - name: terminal
    panes:
      - # Working terminal

options:
  mouse: on
  base-index: 1
  history-limit: 50000
```

## Web Development (Frontend + Backend)

Full-stack web development with separate frontend and backend windows.

```yaml
session:
  name: web-fullstack
  base_dir: ~/projects/web-app

windows:
  - name: editor
    layout: main-vertical
    panes:
      - nvim .
      - # Terminal pane

  - name: frontend
    dir: ./frontend
    panes:
      - cmd: npm run dev
      - cmd: npm run test -- --watch
        split: vertical

  - name: backend
    dir: ./backend
    panes:
      - cmd: npm run dev
      - cmd: npm run test -- --watch
        split: vertical

  - name: database
    panes:
      - cmd: docker-compose up postgres redis
      - cmd: docker-compose logs -f
        split: vertical

  - name: logs
    layout: tiled
    panes:
      - tail -f logs/app.log
      - tail -f logs/error.log
      - tail -f logs/access.log
      - htop

options:
  mouse: on
  base-index: 1
  history-limit: 100000

env:
  NODE_ENV: development
  DATABASE_URL: postgresql://localhost/myapp_dev
  REDIS_URL: redis://localhost:6379
```

## CTF Challenge Development

CTF challenge development with Docker, testing, and notes.

```yaml
session:
  name: ctf-challenge
  base_dir: ~/ctf/challenges/web-001

windows:
  - name: editor
    layout: main-vertical
    panes:
      - nvim challenge.py
      - # Testing pane

  - name: docker
    dir: ./docker
    panes:
      - cmd: docker-compose up
      - cmd: docker-compose logs -f
        split: horizontal

  - name: exploit
    dir: ./exploit
    panes:
      - cmd: python3 solve.py
      - # Manual testing pane
        split: vertical

  - name: monitoring
    layout: tiled
    panes:
      - tcpdump -i any -w capture.pcap
      - docker stats
      - htop
      - # Monitoring pane

  - name: notes
    panes:
      - nvim NOTES.md

options:
  mouse: on
  base-index: 1
  history-limit: 50000

env:
  CHALLENGE_ID: web-001
  FLAG: flag{test_flag}
  DEBUG: "true"
```

## Kubernetes Development

Development with Kubernetes cluster management.

```yaml
session:
  name: k8s-dev
  base_dir: ~/projects/k8s-app

windows:
  - name: editor
    panes:
      - nvim .

  - name: kubectl
    layout: even-vertical
    panes:
      - kubectl get pods -w
      - kubectl get svc -w
      - # Commands pane

  - name: logs
    layout: tiled
    panes:
      - kubectl logs -f deployment/frontend
      - kubectl logs -f deployment/backend
      - kubectl logs -f deployment/worker
      - # Additional logs

  - name: port-forward
    panes:
      - kubectl port-forward svc/frontend 8080:80
      - kubectl port-forward svc/backend 3000:3000

  - name: build
    panes:
      - # Build and deploy pane

options:
  mouse: on
  base-index: 1

env:
  KUBECONFIG: ~/.kube/config
  NAMESPACE: development
```

## Data Science / Jupyter

Data science workflow with Jupyter, database, and analysis.

```yaml
session:
  name: data-science
  base_dir: ~/projects/data-analysis

windows:
  - name: jupyter
    panes:
      - jupyter lab

  - name: database
    panes:
      - docker-compose up postgres
      - psql -h localhost -U postgres
        split: vertical

  - name: scripts
    layout: main-vertical
    panes:
      - nvim analysis.py
      - python3 analysis.py

  - name: terminal
    panes:
      - # Working terminal

options:
  mouse: on
  base-index: 1
  history-limit: 100000

env:
  JUPYTER_PORT: "8888"
  DATABASE_URL: postgresql://localhost/analysis_db
```

## Microservices Development

Development environment for microservices architecture.

```yaml
session:
  name: microservices
  base_dir: ~/projects/microservices

windows:
  - name: editor
    panes:
      - nvim .

  - name: auth-service
    dir: ./services/auth
    panes:
      - npm run dev
      - npm run test:watch
        split: vertical

  - name: user-service
    dir: ./services/user
    panes:
      - npm run dev
      - npm run test:watch
        split: vertical

  - name: payment-service
    dir: ./services/payment
    panes:
      - npm run dev
      - npm run test:watch
        split: vertical

  - name: infrastructure
    layout: tiled
    panes:
      - docker-compose up postgres
      - docker-compose up redis
      - docker-compose up rabbitmq
      - docker-compose logs -f

  - name: monitoring
    layout: tiled
    panes:
      - htop
      - docker stats
      - # Monitoring pane
      - # Monitoring pane

options:
  mouse: on
  base-index: 1
  history-limit: 100000

env:
  NODE_ENV: development
  LOG_LEVEL: debug
```

## DevOps / Infrastructure

Infrastructure management and monitoring.

```yaml
session:
  name: devops
  base_dir: ~/infrastructure

windows:
  - name: terraform
    dir: ./terraform
    panes:
      - # Terraform commands
      - terraform plan
        split: vertical

  - name: ansible
    dir: ./ansible
    panes:
      - # Ansible commands

  - name: monitoring
    layout: tiled
    panes:
      - ssh server1 "htop"
      - ssh server2 "htop"
      - ssh server3 "htop"
      - docker stats

  - name: logs
    layout: tiled
    panes:
      - ssh server1 "tail -f /var/log/app.log"
      - ssh server2 "tail -f /var/log/app.log"
      - ssh server3 "tail -f /var/log/app.log"
      - # Aggregate logs

  - name: kubernetes
    panes:
      - kubectl get pods -w
      - kubectl get nodes -w
        split: vertical

options:
  mouse: on
  base-index: 1
  history-limit: 100000
```

## Go Development

Go project development with testing and debugging.

```yaml
session:
  name: go-project
  base_dir: ~/projects/go-app

windows:
  - name: editor
    layout: main-vertical
    panes:
      - nvim .
      - # Terminal

  - name: run
    panes:
      - go run ./cmd/server
      - # Command pane
        split: vertical

  - name: test
    panes:
      - go test -v ./...
      - go test -bench=. -benchmem ./...
        split: vertical

  - name: tools
    layout: even-vertical
    panes:
      - # golangci-lint run
      - # go vet ./...
      - # Build commands

options:
  mouse: on
  base-index: 1
  history-limit: 50000

env:
  CGO_ENABLED: "0"
  GOOS: linux
```

## Documentation Writing

Documentation writing with live preview.

```yaml
session:
  name: docs
  base_dir: ~/projects/docs

windows:
  - name: editor
    layout: main-vertical
    panes:
      - nvim docs/
      - # Terminal

  - name: preview
    panes:
      - mkdocs serve
      - # Browser/preview commands
        split: vertical

  - name: build
    panes:
      - # Build and deploy commands

options:
  mouse: on
  base-index: 1
```

## Tips for Creating Configurations

### Start Simple

Begin with a basic configuration and add complexity as needed:

```yaml
session:
  name: simple
  base_dir: .

windows:
  - name: main
    panes:
      - # Your main work pane
```

### Use Meaningful Names

Choose descriptive names for sessions and windows:

```yaml
# Good
session:
  name: myapp-frontend-dev

# Not as good
session:
  name: dev
```

### Group Related Panes

Use layouts to organize related panes:

```yaml
windows:
  - name: services
    layout: tiled  # All services visible
    panes:
      - docker-compose up postgres
      - docker-compose up redis
      - docker-compose up rabbitmq
```

### Use Comments

Document your configuration with YAML comments:

```yaml
windows:
  - name: build
    panes:
      - npm run build:watch  # Auto-rebuild on changes
      - # Manual build testing pane
```

### Test Incrementally

Build your configuration step by step:

1. Start with one window
2. Test with `hive launch`
3. Add more windows
4. Test again
5. Add panes and commands
6. Test final setup
