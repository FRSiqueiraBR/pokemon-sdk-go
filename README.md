# pokemon-sdk-go

Lightweight Go SDK scaffold for Pokemon integrations with explicit architecture boundaries.

## Requirements

- Go 1.25+

## Architecture

- Boundary-first layering to minimize coupling.
- `internal/domain`: business entities and typed errors.
- `internal/application`: use cases, ports, mapper, formatters.
- `internal/infra`: HTTP adapters to external APIs.
- `pkg/pokemon`: public service facade and output models.

See `docs/architecture.md` for full boundary rules and extension checklist.

## Project Tree

```text
cmd/examples/pokemon/main.go
internal/domain/
internal/application/pokemon/
internal/infra/pokeapi/
pkg/pokemon/
tests/application/pokemon/
docs/architecture.md
```

## Usage

```go
ctx := context.Background()
service := pokemon.NewService(pokemon.Config{})

summary, err := service.GetPokemon(ctx, "pikachu", pokemon.FormatSummary)
if err != nil {
    // handle error
}

_ = summary
```

## Make Commands

```bash
make help
make test
make run-example
make fmt
make tidy
```

- `make help`: lista os alvos disponíveis.
- `make test`: executa `go test ./...`.
- `make run-example`: roda o exemplo em `cmd/examples/pokemon`.
- `make fmt`: formata todos os arquivos Go do repositório.
- `make tidy`: executa `go mod tidy`.

## Run Tests

```bash
go test ./...
```

## Design Constraints

- SDK consumers interact with a public service facade, not with HTTP clients.
- Application layer does not import infrastructure packages.
- External payloads are never exposed directly to SDK consumers.
- Output format is strategy-based (`summary`, `detailed`) without duplicating integration logic.
