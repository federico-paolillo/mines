# Go Project Template — Extraction Plan

This document describes how to transform the `mines` project into a reusable Go project template.
The template preserves the project's infrastructure patterns (runner, config, CI, linting, hooks, folder structure)
while stripping all domain-specific code (board, game, matchmaking, server, storage, frontend).

The resulting template is **application-type agnostic**: it works equally well for a CLI tool, a web server, a background worker, or any other kind of Go application.

---

## Target Structure

```
cmd/app/main.go                         # Generic entrypoint
internal/
  runner/
    program.go                          # ProgramE type, StatusCode (renamed from programe.go)
    execution.go                        # Signal handling, goroutine orchestration
    many.go                             # RunMany orchestrator (creates logger, loads config, builds deps)
    many_test.go                        # Tests
  example/
    example.go                          # Placeholder internal package
    example_test.go
  testutils/
    slogt/slogt.go                      # Generic slog-to-testing bridge (kept as-is)
pkg/
  app/
    app.go                              # Minimal composition root
    app_test.go
    config/
      config.go                         # Generic Root config struct with Default()
      load.go                           # configuro loader with "APP" env prefix
      load_test.go
  example/
    example.go                          # Placeholder public package
    example_test.go
.github/workflows/ci.yml               # Go-only CI pipeline
.golangci.yml                           # Linter config (Go 1.26)
.gitignore                              # Simplified (no frontend artifacts)
codecov.yml                             # Go-only coverage config
config.example.yml                      # Generic example configuration
lefthook.yml                            # Go-only hooks and tasks
go.mod                                  # Minimal dependencies
```

---

## Action Checklist

### Phase 1 — Remove Domain-Specific Code

Delete these directories and files entirely:

- [ ] `cmd/server/` — replaced by `cmd/app/`
- [ ] `internal/server/` — HTTP server, handlers, routes, validators, middlewares, `req/`, `res/`
- [ ] `internal/generators/` — RNG board generator
- [ ] `internal/id/` — ID generation
- [ ] `internal/storage/` — storage abstraction and implementations (memory, memcached)
- [ ] `internal/testutils/generator.go` — domain test fixture
- [ ] `internal/testutils/matches.go` — domain test fixture
- [ ] `internal/testutils/requests.go` — domain test fixture
- [ ] `internal/testutils/server.go` — domain test fixture
- [ ] `internal/testutils/printers/` — ASCII board printer
- [ ] `pkg/board/` — board logic
- [ ] `pkg/dimensions/` — location/size types
- [ ] `pkg/game/` — game state machine
- [ ] `pkg/matchmaking/` — match management
- [ ] `pkg/mines/` — domain-specific composition root (replaced by `pkg/app/`)

Delete frontend and related tooling:

- [ ] `fe/` — React frontend
- [ ] `bruno/` — API testing collection
- [ ] `.gemini/` — AI tool config
- [ ] `biome.json` — frontend linter
- [ ] `package.json`, `package-lock.json` — npm
- [ ] `docker-compose.yml` — memcached for local dev

### Phase 2 — Modify Existing Files

#### `internal/runner/programe.go` → Rename to `internal/runner/program.go`

**Why:** The filename `programe.go` appears to be a typo for `program.go`.

The type name `ProgramE` is kept — it follows the Cobra `Command.RunE` convention (a program that returns an error).

Current code:

```go
type ProgramE = func(
    context.Context,
    *mines.Mines,
    *config.Root,
) error

type StatusCode = string

const (
    Ok    StatusCode = "ok"
    NotOk StatusCode = "nok"
)
```

Target code:

```go
package runner

import (
    "context"

    "github.com/yourorg/yourproject/pkg/app"
    "github.com/yourorg/yourproject/pkg/app/config"
)

// ProgramE is a unit of work managed by the runner.
// Programs receive a context (cancelled on shutdown signals), the composition
// root, and the configuration. The naming follows the Cobra RunE convention.
type ProgramE func(context.Context, *app.App, *config.Root) error

type StatusCode string

const (
    Ok    StatusCode = "ok"
    NotOk StatusCode = "nok"
)
```

Changes:
- [ ] Rename file from `programe.go` to `program.go`
- [ ] Change from type alias (`=`) to defined type for both `ProgramE` and `StatusCode` — this provides type safety (a bare `string` can no longer be passed as `StatusCode`, a bare `func` can no longer be passed as `ProgramE`)
- [ ] Replace `*mines.Mines` with `*app.App` (generalized composition root)

#### `internal/runner/execution.go`

Current issues:
1. `signal.Notify(sigtermChan, os.Interrupt)` only handles `os.Interrupt`, not `SIGTERM`. Containers (Docker, Kubernetes) and systemd send `SIGTERM` for graceful shutdown.
2. Variable name `sigtermChan` is misleading since it only captures interrupts.
3. Function signatures reference `*mines.Mines` instead of the generalized `*app.App`.

Changes:
- [ ] Add `syscall.SIGTERM` to `signal.Notify`: `signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)`
- [ ] Rename `sigtermChan` to `signalChan`
- [ ] Replace `*mines.Mines` with `*app.App` in `runManyPrograms` and `runOneProgram` signatures
- [ ] Add `import "syscall"`
- [ ] Update `runOneProgram` to call `program(ctx, appRoot, cfg)` instead of `program(ctx, mines, cfg)`

Target `runManyPrograms` signature:

```go
func runManyPrograms(
    ctx context.Context,
    logger *slog.Logger,
    appRoot *app.App,
    cfg *config.Root,
    programs ...ProgramE,
) StatusCode
```

Target `runOneProgram` signature:

```go
func runOneProgram(
    ctx context.Context,
    wg *sync.WaitGroup,
    errChan chan<- error,
    appRoot *app.App,
    cfg *config.Root,
    program ProgramE,
)
```

#### `internal/runner/many.go`

The runner owns logger creation, config loading, and composition root construction. This is intentional:
the runner will eventually create an OpenTelemetry-bound logger and inject tracing via context, so it must
control the lifecycle of these cross-cutting concerns. This also guarantees one logger per execution.

Changes:
- [ ] Replace `mines.NewMines` with `app.NewApp` (generalized composition root)
- [ ] Replace imports of `mines` package with `app` package
- [ ] Keep logger creation, config loading, and dep construction inside `RunMany`
- [ ] Pass `*app.App` instead of `*mines.Mines` to `runManyPrograms`

Target code:

```go
package runner

import (
    "context"
    "log/slog"
    "os"
    "time"

    "github.com/yourorg/yourproject/pkg/app"
    "github.com/yourorg/yourproject/pkg/app/config"
)

func RunMany(
    ctx context.Context,
    programs ...ProgramE,
) StatusCode {
    logger := slog.New(
        slog.NewJSONHandler(os.Stdout, nil),
    )

    cfg, err := config.Load()
    if err != nil {
        logger.Error(
            "runner: failed to load configuration",
            slog.Any("err", err),
        )

        return NotOk
    }

    appRoot := app.NewApp(logger, cfg)

    startTime := time.Now()

    statusCode := runManyPrograms(ctx, logger, appRoot, cfg, programs...)

    runtimeDuration := time.Since(startTime)

    logger.Info(
        "runner: program has completed. This does not indicate success",
        slog.Float64("runtime_s", runtimeDuration.Seconds()),
        slog.Any("status_code", statusCode),
    )

    return statusCode
}
```

Also fix: `starTime` typo on current line 56 → `startTime`. Use `time.Since(startTime)` instead of manual `endTime.Sub(starTime)`.

Note: `NewApp` receives the runner's logger so that all application components share the same logger instance (important for future OTel integration). `NewApp` does not return an error because it currently cannot fail — add an error return only when initialization can genuinely fail.

#### `internal/runner/many_test.go`

- [ ] Update all test program functions from `func(_ context.Context, _ *mines.Mines, _ *config.Root) error` to `func(_ context.Context, _ *app.App, _ *config.Root) error`
- [ ] Replace imports of `mines` package with `app` package
- [ ] Remove `t.Helper()` from program closures (not needed — these are not test helpers, they are test subjects)

#### `.golangci.yml`

- [ ] Change `go: "1.24"` to `go: "1.26"`

#### `.github/workflows/ci.yml`

- [ ] Remove "Setup Node" step
- [ ] Remove "Setup kiota" step
- [ ] Remove "npm ci" step
- [ ] Remove "Swagger" step
- [ ] Remove "Frontend Coverage" step
- [ ] Update `go-version` from `"1.25"` to `"1.26"`
- [ ] In Codecov upload, change `files: ./coverage.txt,./fe/coverage/lcov.info` to `files: ./coverage.txt`

Target workflow:

```yaml
name: build

on:
  push:
    branches: ["main"]

jobs:
  build:
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: "1.26"

      - name: Build
        run: go tool lefthook run build

      - name: Test
        run: go tool lefthook run test

      - name: Vet
        run: go tool lefthook run lint

      - name: Coverage
        run: go tool lefthook run cover

      - name: Upload coverage reports to Codecov
        uses: codecov/codecov-action@v5.5.2
        with:
          token: ${{ secrets.CODECOV_TOKEN }}
          files: ./coverage.txt
          fail_ci_if_error: true
          disable_search: true
```

#### `lefthook.yml`

Remove all frontend-related commands and the `swagger` task entirely.

Target:

```yaml
lint:
  parallel: true
  commands:
    golangci:
      run: go tool golangci-lint run

test:
  parallel: true
  commands:
    gotest:
      run: go test -race -shuffle=on ./...

format:
  commands:
    golangci:
      run: go tool golangci-lint run --fix

cover:
  commands:
    gotest:
      run: go test -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic ./...

build:
  commands:
    gobuild:
      run: go build -v ./...

pre-push:
  commands:
    test:
      glob:
        - "*.go"
      run: go tool lefthook run test

pre-commit:
  commands:
    lint:
      glob:
        - "*.go"
      run: go tool lefthook run lint
```

#### `codecov.yml`

- [ ] Remove `internal/gc` from ignore list (does not exist in template)

Target:

```yaml
coverage:
  precision: 0
  round: up
ignore:
  - internal/testutils
  - cmd/
  - program.go
parsers:
  go:
    partials_as_hits: true
```

#### `.gitignore`

- [ ] Remove `node_modules/`, `dist/`, `docs/swagger.yaml`, `fe/coverage/`

Target:

```gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Go workspace file
go.work
go.work.sum

# env file
.env

# coverage
coverage.txt

config.yml

.DS_Store
```

#### `config.example.yml`

Replace domain-specific config with a generic placeholder:

```yaml
app:
  name: myapp
  debug: false
```

#### `go.mod`

- [ ] Change module path to `github.com/yourorg/yourproject` (placeholder)
- [ ] Change `go 1.25` to `go 1.26`
- [ ] Remove domain-specific dependencies: `gin`, `gomemcache`, `httpexpect`, `testcontainers`, `testcontainers-go/modules/memcached`, `go-playground/validator`
- [ ] Remove `swag` from `tool` directive
- [ ] Keep: `configuro`, `stretchr/testify`, `lefthook`, `golangci-lint`
- [ ] Run `go mod tidy` to clean up `go.sum`

### Phase 3 — Create New Files

#### `cmd/app/main.go`

The entrypoint simply registers `ProgramE` functions and delegates to `RunMany`.
Config loading, logger creation, and dependency construction happen inside the runner.

```go
package main

import (
    "context"
    "os"

    "github.com/yourorg/yourproject/internal/runner"
    "github.com/yourorg/yourproject/pkg/app"
    "github.com/yourorg/yourproject/pkg/app/config"
)

// ExampleProgram demonstrates a ProgramE that uses the composition root.
// Replace this with your actual program logic.
func ExampleProgram(
    ctx context.Context,
    appRoot *app.App,
    cfg *config.Root,
) error {
    appRoot.Logger.Info(
        "example program started",
        "app_name", cfg.App.Name,
        "debug", cfg.App.Debug,
    )

    // Do your work here. For long-running programs (servers, workers),
    // block until ctx is cancelled. For one-shot programs (CLI commands,
    // migrations), just do the work and return.

    <-ctx.Done()

    appRoot.Logger.Info("example program shutting down")

    return nil
}

func main() {
    statusCode := runner.RunMany(
        context.Background(),
        ExampleProgram,
    )

    if statusCode == runner.NotOk {
        os.Exit(1)
    }
}
```

Note: The `ProgramE` receives the composition root and config from the runner. This ensures the logger (and future OTel tracing) is consistently wired across all programs.

#### `pkg/app/app.go`

The composition root receives the logger from the runner (not creating its own).
This ensures the logger is the single instance managed by the runner, ready for
future OpenTelemetry integration.

```go
package app

import (
    "log/slog"

    "github.com/yourorg/yourproject/pkg/app/config"
)

// App is the composition root. Add your application's dependencies as fields.
type App struct {
    Logger *slog.Logger
    Config *config.Root
}

func NewApp(logger *slog.Logger, cfg *config.Root) *App {
    return &App{
        Logger: logger,
        Config: cfg,
    }
}
```

Note: `NewApp` does not return an error because it currently cannot fail. Add an error return only when initialization can genuinely fail (e.g., opening a database connection). Do not return errors that are always nil.

#### `pkg/app/app_test.go`

```go
package app_test

import (
    "log/slog"
    "os"
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/yourorg/yourproject/pkg/app"
    "github.com/yourorg/yourproject/pkg/app/config"
)

func TestNewAppReturnsApp(t *testing.T) {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    cfg := config.Default()

    application := app.NewApp(logger, cfg)

    require.NotNil(t, application)
    require.Equal(t, logger, application.Logger)
    require.Equal(t, cfg, application.Config)
}
```

#### `pkg/app/config/config.go`

```go
package config

// Root is the top-level configuration structure.
// Add your application's configuration sections as fields.
type Root struct {
    App App
}

type App struct {
    Name  string
    Debug bool
}

func Default() *Root {
    return &Root{
        App: App{
            Name:  "myapp",
            Debug: false,
        },
    }
}
```

#### `pkg/app/config/load.go`

```go
package config

import (
    "fmt"

    "github.com/sherifabdlnaby/configuro"
)

func Load() (*Root, error) {
    cfguro, err := configuro.NewConfig(
        configuro.WithLoadFromEnvVars("APP"),
        configuro.WithLoadFromConfigFile("config.yml", false),
        configuro.WithoutEnvConfigPathOverload(),
        configuro.WithoutLoadDotEnv(),
    )
    if err != nil {
        //nolint:errorlint // We do not want to wrap and leak errors that are not under our control
        return nil, fmt.Errorf(
            "config: failed to setup configuro. %v",
            err,
        )
    }

    cfg := Default()

    err = cfguro.Load(cfg)
    if err != nil {
        //nolint:errorlint // We do not want to wrap and leak errors that are not under our control
        return nil, fmt.Errorf(
            "config: failed to bind configuration. %v",
            err,
        )
    }

    return cfg, nil
}
```

#### `pkg/app/config/load_test.go`

```go
package config_test

import (
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/yourorg/yourproject/pkg/app/config"
)

func TestConfigurationLoadsFromEnvVars(t *testing.T) {
    t.Setenv("APP_APP_NAME", "testapp")
    t.Setenv("APP_APP_DEBUG", "true")

    cfg, err := config.Load()

    require.NoError(t, err)
    require.Equal(t, "testapp", cfg.App.Name)
    require.True(t, cfg.App.Debug)
}
```

#### `pkg/example/example.go`

```go
// Package example is a placeholder public package.
// Replace this with your domain logic.
// Public packages in pkg/ are importable by external consumers of this module.
package example

func Hello() string {
    return "hello from pkg/example"
}
```

#### `pkg/example/example_test.go`

```go
package example_test

import (
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/yourorg/yourproject/pkg/example"
)

func TestHello(t *testing.T) {
    require.Equal(t, "hello from pkg/example", example.Hello())
}
```

#### `internal/example/example.go`

```go
// Package example is a placeholder internal package.
// Replace this with your internal implementation details.
// Internal packages cannot be imported by external consumers of this module.
package example

func Greet(name string) string {
    return "hello, " + name
}
```

#### `internal/example/example_test.go`

```go
package example_test

import (
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/yourorg/yourproject/internal/example"
)

func TestGreet(t *testing.T) {
    require.Equal(t, "hello, world", example.Greet("world"))
}
```

---

## Bugs and Idiomacy Fixes

### Bug: Data Race in `internal/server/program.go:50-65`

This file is removed in the template, but the pattern is documented here as an anti-pattern to avoid when writing `ProgramE` functions.

The current code starts `ListenAndServe` in a goroutine and writes to `err` (line 51), then reads `err` on return (line 65) without synchronization — a data race. Additionally, `server.Shutdown(ctx)` on line 63 uses the same context that triggered `ctx.Done()`, giving the shutdown zero grace period.

**Correct pattern for long-running programs:**

```go
func MyServerProgram(ctx context.Context, appRoot *app.App, cfg *config.Root) error {
    server := &http.Server{Addr: ":8080"}

    errCh := make(chan error, 1)
    go func() {
        if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            errCh <- err
        }
        close(errCh)
    }()

    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := server.Shutdown(shutdownCtx); err != nil {
        return fmt.Errorf("server shutdown failed: %w", err)
    }

    // Check if ListenAndServe returned a real error
    if err := <-errCh; err != nil {
        return fmt.Errorf("server failed: %w", err)
    }

    return nil
}
```

Key points:
- Use a channel to communicate errors from the listener goroutine (no shared variable)
- Create a **fresh context with timeout** for `Shutdown()` — do not reuse the cancelled context
- Check for `http.ErrServerClosed` which is expected after `Shutdown()`

### Bug: Missing SIGTERM Signal (`internal/runner/execution.go:29`)

Only `os.Interrupt` is registered. Containers (Docker, Kubernetes) and systemd send `SIGTERM` for graceful shutdown. The template adds `syscall.SIGTERM`.

### Typo: `programe.go` Filename

Renamed to `program.go`.

### Type Safety: Type Aliases vs Defined Types

`ProgramE` and `StatusCode` use type aliases (`=`), which provide no type safety — any `string` can be used where `StatusCode` is expected, and any matching `func` can be passed as `ProgramE`. The template uses defined types instead.

### Phantom Error Return: `pkg/mines/mines.go:28`

`NewMines` returns `(*Mines, error)` but the error is always `nil`. The template's `NewApp` returns `*App` without an error. **Rule: do not return errors that can never occur.** Add an error return only when initialization can genuinely fail.

### Runner Owns the Composition Root (Intentional)

`RunMany` loads config and constructs the composition root. This is by design: the runner creates the logger and will eventually bind it to OpenTelemetry, so it must control the lifecycle of cross-cutting concerns. This guarantees one logger per execution and a consistent tracing setup across all programs. The template generalizes `*mines.Mines` to `*app.App` but preserves the same ownership pattern.

### Minor: `starTime` Typo (`internal/runner/many.go:56`)

Rename to `startTime`. Also simplify: use `time.Since(startTime)` instead of `endTime.Sub(starTime)`.

---

## Modernization Notes (Go 1.26)

### Version Bumps

Update all Go version references to 1.26:

| File | Field | From | To |
|------|-------|------|----|
| `go.mod` | `go` directive | `1.25` | `1.26` |
| `.golangci.yml` | `run.go` | `"1.24"` | `"1.26"` |
| `.github/workflows/ci.yml` | `go-version` | `"1.25"` | `"1.26"` |

### Already-Modern Patterns

The codebase already uses several modern Go patterns that should be preserved:

- **Range over integers** (Go 1.22): `for x := range size.Width` — used throughout `pkg/board/`
- **`log/slog`** (Go 1.21): structured logging via the standard library
- **`go tool` directives** (Go 1.24): `tool` block in `go.mod` for lefthook, golangci-lint
- **`t.Context()`** (Go 1.24): used in `many_test.go` instead of manual context creation
- **`crypto/rand.Text()`** (Go 1.24): used in `internal/id/` for secure random strings

### Recommendations for Template Users

- Use `slog.With()` to create scoped loggers with default attributes (e.g., per-component or per-request)
- Consider `context.AfterFunc` (Go 1.21) for registering cleanup callbacks on context cancellation
- The `configuro` library works but has not been actively maintained since ~2020. If you need more features or long-term maintenance, consider:
  - [`koanf`](https://github.com/knadh/koanf) — lightweight, composable, actively maintained
  - [`viper`](https://github.com/spf13/viper) — most popular, feature-rich, heavier dependency tree

---

## Verification

After completing all changes, verify the template works:

```bash
# Must compile
go build -v ./...

# All tests pass with race detection
go test -race -shuffle=on ./...

# Linter passes
go tool golangci-lint run

# Lefthook tasks work
go tool lefthook run lint
go tool lefthook run test
go tool lefthook run cover
go tool lefthook run build
```

---

## Quick Start (for template users)

After cloning the template:

1. Replace the module path: `github.com/yourorg/yourproject` → your actual module path
2. Rename the `APP` env var prefix in `pkg/app/config/load.go` to match your project
3. Update `config.example.yml` and `pkg/app/config/config.go` with your configuration fields
4. Add your dependencies to the `App` struct in `pkg/app/app.go` and wire them in `NewApp`
5. Replace `pkg/example/` and `internal/example/` with your actual packages
6. Replace `ExampleProgram` in `cmd/app/main.go` with your actual `ProgramE` function(s)
7. Run `go mod tidy`
