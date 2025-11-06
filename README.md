# PROJECT_NAME

[![Go Build & Docker Build](https://github.com/OWNER/PROJECT_NAME/actions/workflows/build.yaml/badge.svg)](https://github.com/OWNER/PROJECT_NAME/actions/workflows/build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/OWNER/PROJECT_NAME)](https://goreportcard.com/report/github.com/OWNER/PROJECT_NAME)
[![Go Doc](https://godoc.org/github.com/OWNER/PROJECT_NAME?status.svg)](https://godoc.org/github.com/OWNER/PROJECT_NAME)

> A well-structured Go application template

PROJECT_NAME is a production-ready Go application template with clean architecture, Kubernetes integration, modern CLI interfaces, and robust observability.

## Features

- **Clean Architecture** - Clear separation of concerns with `cmd/`, `pkg/`, and `internal/` structure
- **CLI Support** - Modern CLI built with Cobra and Bubbletea
- **Kubernetes Ready** - Operator framework integration with controller-runtime
- **Observability** - OpenTelemetry tracing, Prometheus metrics, structured logging
- **Testing** - Comprehensive test structure with mocks and helpers
- **CI/CD** - GitHub Actions workflows for build, test, and release

## Quick Start

```bash
# Clone this template
git clone https://github.com/OWNER/PROJECT_NAME.git
cd PROJECT_NAME

# Run the initialization script
./scripts/init.sh YOUR_PROJECT_NAME YOUR_GITHUB_OWNER

# Install dependencies
go mod tidy

# Build
make build

# Run tests
make test
```

## Repository Structure

```text
.
├── api/                          # Kubernetes Custom Resource Definitions
│   └── v1alpha1/                # API version
├── cmd/                         # Application entry points
│   ├── cli/                     # CLI application
│   └── server/                  # Server application
├── config/                      # Kubernetes manifests and configuration
│   ├── crd/                    # Custom Resource Definitions
│   └── rbac/                   # Role-Based Access Control
├── internal/                    # Private application code
│   ├── cli/                    # CLI-specific utilities
│   └── utils/                  # Shared utilities
├── pkg/                        # Public library code
│   ├── client/                 # Client abstractions
│   ├── middleware/             # HTTP middleware
│   ├── models/                 # Data models and DTOs
│   └── service/                # Business logic services
├── hack/                       # Development utilities
├── scripts/                    # Utility scripts
├── go.mod                      # Go module definition
├── Makefile                    # Build automation
└── README.md                   # Project documentation
```

## Development

### Prerequisites

- Go 1.21+
- Make
- Docker (for container builds)

### Building

```bash
# Build all binaries
make build

# Run tests
make test

# Run linter
make lint

# Generate code (CRDs, swagger, etc.)
make generate
```

### Running Locally

```bash
# Run CLI
./bin/PROJECT_NAME-cli --help

# Run server
./bin/PROJECT_NAME-server
```

## Configuration

Configuration is managed through:

- Default values in code
- Config file (YAML/JSON)
- Environment variables
- Command-line flags

See `config.yaml` for example configuration.

## Documentation

- **[Project Structure Guide](PROJECT_STRUCTURE_GUIDE.md)** - Comprehensive guide to the project structure and conventions

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request
