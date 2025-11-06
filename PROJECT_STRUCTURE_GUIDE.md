# PROJECT_NAME Project Structure & Conventions Guide

This document provides a comprehensive guide to the PROJECT_NAME project structure, conventions, and framework choices. Use this as a template for bootstrapping similar Go projects with the same architectural patterns.

## Project Overview

PROJECT_NAME is a well-structured Go application template that follows clean architecture principles with clear separation of concerns.

## Repository Structure

```text
.
├── api/                          # Kubernetes Custom Resource Definitions
│   └── v1alpha1/                # API version
│       ├── groupversion_info.go # API group metadata
│       └── zz_generated.deepcopy.go # Generated code
├── cmd/                         # Application entry points
│   ├── cli/                     # CLI application
│   │   ├── main.go             # CLI main entry point
│   │   └── cmd_*.go            # Individual command implementations
│   └── server/                  # Server application
│       ├── main.go             # Server main entry point
│       └── serve.go            # Server command implementation
├── config/                      # Kubernetes manifests and configuration
│   ├── crd/                    # Custom Resource Definitions
│   └── rbac/                   # Role-Based Access Control
├── internal/                    # Private application code
│   ├── cli/                    # CLI-specific utilities
│   │   ├── async_operation/    # Async operation handling
│   │   ├── cmd/                # Command definitions and flags
│   │   └── pretty_print/       # CLI output formatting
│   └── utils/                  # Shared utilities
│       ├── gin.go              # Gin HTTP framework utilities
│       ├── interrupt_handler.go # Signal handling
│       └── o11y.go             # Observability utilities
├── pkg/                        # Public library code
│   ├── client/                 # Client abstractions
│   ├── middleware/             # HTTP middleware
│   │   └── auth/               # Authentication middleware
│   ├── models/                 # Data models and DTOs
│   ├── operator/               # Kubernetes operator logic
│   ├── service/                # Business logic services
│   │   ├── api/                # HTTP API handlers
│   │   └── models/             # Service-specific models
│   └── swagger/                # API documentation
├── hack/                       # Development utilities
│   └── boilerplate.go.txt      # Code generation templates
├── scripts/                    # Utility scripts
│   └── init.sh                 # Project initialization script
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── Makefile                    # Build automation
├── config.yaml                 # Default configuration
├── Dockerfile                  # Container image definition
└── README.md                   # Project documentation
```

## Framework & Technology Stack

### Core Frameworks

1. **Go 1.21+** - Primary programming language
2. **Cobra** - CLI framework for command-line applications
3. **Viper** - Configuration management
4. **Gin** - HTTP web framework
5. **Controller Runtime** - Kubernetes operator framework

### Key Dependencies

- **HTTP & API**:
  - `github.com/gin-gonic/gin` - HTTP web framework
  - `github.com/swaggo/gin-swagger` - API documentation
  - `github.com/swaggo/swag` - Swagger code generation

- **Observability**:
  - `go.opentelemetry.io/otel` - OpenTelemetry tracing
  - `go.uber.org/zap` - Structured logging
  - `github.com/prometheus/client_golang` - Metrics collection

- **Kubernetes**:
  - `k8s.io/client-go` - Kubernetes client library
  - `sigs.k8s.io/controller-runtime` - Operator framework
  - `k8s.io/apimachinery` - Kubernetes API machinery

- **CLI & UX**:
  - `github.com/charmbracelet/bubbletea` - TUI framework
  - `github.com/charmbracelet/lipgloss` - Styling for TUI
  - `github.com/spf13/cobra` - CLI framework

## Architectural Patterns

### 1. Clean Architecture

The project follows clean architecture principles with clear layer separation:

- **API Layer** (`api/`): Kubernetes CRDs and API definitions
- **Application Layer** (`cmd/`): Entry points and application orchestration
- **Domain Layer** (`pkg/service/`): Business logic and domain models
- **Infrastructure Layer** (`pkg/client/`): External integrations

### 2. Command Pattern

CLI commands are organized using the command pattern:

```go
// cmd/cli/main.go
cmdRoot := cmd.NewCliRootCmd()
cmdRoot.AddCommand(cmdGenerate)
cmdRoot.AddCommand(cmdGet)
cmdRoot.AddCommand(cmdSet)
```

### 3. Middleware Pattern

HTTP middleware is implemented for cross-cutting concerns:

```go
// pkg/middleware/auth/gin.go
type ginAuthMiddleware struct {
    // Middleware configuration
}
```

### 4. Option Pattern

Configuration is handled using the functional options pattern:

```go
// pkg/service/api/server.go
func NewServer(opts ...Option) *Server {
    server := &Server{...}
    for _, opt := range opts {
        opt(server)
    }
    return server
}
```

## Code Conventions

### 1. Package Organization

- **`cmd/`**: Application entry points only
- **`internal/`**: Private application code (not importable)
- **`pkg/`**: Public library code (importable by other projects)
- **`api/`**: Kubernetes API definitions

### 2. Naming Conventions

- **Packages**: Lowercase, descriptive names (`pretty_print`, `client`)
- **Types**: PascalCase (`Server`, `UserResponse`)
- **Functions**: PascalCase for public, camelCase for private
- **Variables**: camelCase
- **Constants**: PascalCase or UPPER_CASE

### 3. Error Handling

Uses custom error types for better user experience:

```go
// pkg/models/api_error.go
type APIError struct {
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}
```

### 4. Configuration Management

Configuration follows a hierarchical approach:

1. **Default values** in code
2. **Config file** (YAML/JSON)
3. **Environment variables**
4. **Command-line flags**

### 5. Testing Conventions

- Test files end with `_test.go`
- Table-driven tests for comprehensive coverage
- Helper functions use `t.Helper()`
- Mock implementations in `mock/` subdirectories

## Build & Development

### 1. Makefile Targets

```makefile
.PHONY: build
build: generate
  goreleaser build --clean --snapshot

.PHONY: generate
generate: controller-gen swag
  controller-gen object:headerFile="./hack/boilerplate.go.txt" paths="./..."

.PHONY: test
test: lint
  go test -race ./... -coverprofile=coverage.out
```

### 2. Code Generation

- **Controller-gen**: Generates Kubernetes CRDs and RBAC
- **Swag**: Generates API documentation
- **Deepcopy**: Generates deep copy methods for CRDs

### 3. Observability

- **OpenTelemetry**: Distributed tracing
- **Prometheus**: Metrics collection
- **Zap**: Structured logging

## Development Workflow

### 1. Project Setup

```bash
# Initialize Go module
go mod init github.com/your-org/your-project

# Install dependencies
go mod tidy

# Generate code
make generate

# Run tests
make test
```

### 2. Adding New Commands

1. Create command file in `cmd/cli/`
2. Define command structure with Cobra
3. Add to main command tree in `main.go`
4. Implement business logic in `pkg/service/`

### 3. Adding New API Endpoints

1. Define handler in `pkg/service/api/`
2. Add route registration
3. Update Swagger documentation
4. Add tests for the endpoint

### 4. Adding New Kubernetes Resources

1. Define CRD in `api/v1alpha1/`
2. Generate deepcopy methods
3. Implement controller logic in `pkg/operator/`
4. Update RBAC in `config/rbac/`

## Best Practices

### 1. Error Handling

- Use custom error types for better UX
- Provide actionable error messages
- Log errors with appropriate levels
- Return errors, don't panic

### 2. Configuration

- Provide sensible defaults
- Support multiple configuration sources
- Validate configuration on startup
- Use structured configuration types

### 3. Testing

- Write table-driven tests
- Use mocks for external dependencies
- Test error conditions
- Aim for high test coverage

### 4. Documentation

- Document public APIs
- Use meaningful variable names
- Add examples in code comments
- Keep README up to date

### 5. Security

- Validate all inputs
- Use secure defaults
- Implement proper authentication
- Log security events

This structure provides a solid foundation for building production-ready Go applications with Kubernetes integration, modern CLI interfaces, and robust observability.
