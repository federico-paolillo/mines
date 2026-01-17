# Project: Mines (Minesweeper Clone)

## 1. Project Overview

A client-server Minesweeper game designed for scalability and concurrency.

- **Backend**: Go 1.25 HTTP API with OpenAPI specification.
- **Frontend**: Preact (Vite) Single Page Application.
- **Key Features**:
  - **Stateless/Scalable**: Game state stored in Memcached with a 2-hour TTL.
  - **Concurrency Control**: Optimistic Offline Lock (using Memcached CAS) to handle concurrent moves.
  - **Durability**: Non-goal; games are transient.

## 2. Tech Stack & Tools

### Backend

- **Language**: Go 1.25
- **Web Framework**: Gin (`github.com/gin-gonic/gin`).
- **Storage**:
  - **In-memory**: For local development/testing.
  - **Memcached**: For production/staging (Scalability + TTL + CAS).
- **Linting**: `golangci-lint` (config: `.golangci.yml`).
- **Testing**: Native `testing`, `testify` (assertions), `Testcontainers` (integration).
- **Documentation**: Swagger/OpenAPI (generated via `swag`).

### Frontend

- **Framework**: Preact (with `preact-iso` for routing).
- **Build Tool**: Vite.
- **Styling**: Tailwind CSS.
- **Linting/Formatting**: Biome (config: `biome.json`).
- **API Client**: TypeScript client generated via `Kiota`.

### Automation & DevOps

- **Task Runner**: Lefthook (`lefthook.yml`).
- **Containerization**: Docker, Docker Compose (Memcached), Podman (Mac integration).
- **CLI**: Gemini CLI (scoped via root `package.json`).

## 3. Architecture & Patterns

### Backend

- **Project Structure** (Standard Go Layout):
  - `cmd/server`: Main entry point.
  - `internal/`: Private application logic (handlers, storage implementations, middleware).
  - `pkg/`: Public domain logic (Minesweeper rules, board generation).
- **Composition Root** (`pkg/mines/mines.go`):
  - Acts as a dependency injection container.
  - Passed around like `context.Context` to provide access to services/config without global state.
- **Concurrency Model**:
  - Optimistic locking ensures data integrity when multiple clients interact with the same game.
  - Relies on Memcached's CAS (Check-And-Set) versioning.

### Frontend

- **Structure** (`fe/`):
  - `pages/`: Top-level routable components.
  - `components/`: Reusable UI widgets.
  - `client/`: **Auto-generated** API client. **Do not modify manually.**
- **Routing**: `preact-iso` handles client-side routing.

## 4. Development Workflow

### Prerequisites

- Go 1.25+
- Node.js v21.x
- Docker & Docker Compose
- Podman (Recommended for macOS Testcontainers support)

### Setup

1.  **Configure Backend**: `cp config.example.yml config.yml`
2.  **Configure Frontend**: `cp fe/.env.example fe/.env`
3.  **Start Dependencies**: `docker compose up -d` (Starts Memcached)

### Common Commands (Lefthook)

Use `go tool lefthook run <task>` to execute workflows defined in `lefthook.yml`.

| Task        | Command                        | Description                                           |
| :---------- | :----------------------------- | :---------------------------------------------------- |
| **Lint**    | `go tool lefthook run lint`    | Runs `golangci-lint` (Go) and `biome` (TSX).          |
| **Test**    | `go tool lefthook run test`    | Runs Go tests (with race detection) and Vitest.       |
| **Format**  | `go tool lefthook run format`  | Auto-fixes lint errors and formats code.              |
| **Build**   | `go tool lefthook run build`   | Compiles Go binary and builds Vite frontend.          |
| **CodeGen** | `go tool lefthook run swagger` | **Crucial**: Regenerates Swagger docs & Kiota client. |

### Running the App

- **Backend**: `go run ./cmd/server`
- **Frontend**: `cd fe && npm run dev`

## 5. Guidelines for Agents

### Modifying the Backend

1.  **Dependency Injection**: If adding a new service, register it in `pkg/mines/mines.go`.
2.  **API Changes**:
    - Update `routes.go` or handlers (`internal/server/handlers`).
    - Add/Update Swagger comments (See `swag` documentation).
    - **MUST RUN**: `go tool lefthook run swagger` to update the spec and frontend client.
3.  **Storage**:
    - Implement changes in both `internal/storage/memory` AND `internal/storage/memcached`.
    - Verify with `go tool lefthook run test`.

### Modifying the Frontend

1.  **State Management**: Use Preact hooks.
2.  **API Integration**:
    - Use the generated client in `fe/src/client`.
    - If the client is missing a method, check if Backend API changes were generated (`lefthook run swagger`).
3.  **Components**: Follow the `pages` vs `components` split.

### Testing

- **Integration Tests**: Require a running container environment (Podman/Docker).
- **Unit Tests**: `testify` for Go, `vitest` for Preact.
- Avoid needless mocks. Strive to use real implementations. 
- Read `test-setup.ts` for `vitest` tests and use it to avoid repeating common setup and teardown.

## 6. Key Files

- `config.example.yml`: App configuration schema.
- `pkg/mines/mines.go`: **Composition Root** (DI Container).
- `internal/storage/store.go`: Storage interface definition.
- `fe/src/client/`: Generated API client code.

## 7. Configuration Reference

### Backend (`config.yml`)

The backend is configured via a YAML file (parsed at startup).

- **`seed`**: (int) Initial seed for random number generation.
- **`ttl`**: (int) Game state Time-To-Live in hours.
- **`server`**:
  - **`host`**: (string) Hostname to bind to.
  - **`port`**: (int) Port to listen on (Default: `65000`).
- **`memcached`**:
  - **`enabled`**: (bool) Toggle between in-memory (false) and Memcached (true) storage.
  - **`servers`**: (list<string>) List of Memcached server addresses (e.g., `localhost:60001`).

### Frontend (`fe/.env`)

The frontend uses Vite's environment variable system.

- **`VITE_MINESWEEPER_API_URL`**: (url) Base URL for API calls.
  - **Development**: Typically `http://localhost:5173/api` (Vite dev server).
  - **Proxying**: `fe/vite.config.ts` proxies `/api` requests to `http://localhost:65000` (stripping the `/api` prefix) to resolve CORS and path matching during development.
