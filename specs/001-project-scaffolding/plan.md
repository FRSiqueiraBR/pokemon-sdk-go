# Implementation Plan: Project Scaffolding

**Branch**: `[001-project-scaffolding]` | **Date**: 2026-02-18 | **Spec**: `specs/001-project-scaffolding/spec.md`
**Input**: Feature specification from `/specs/001-project-scaffolding/spec.md`

## Summary

Criar baseline arquitetural em Go para SDK de integrações HTTP externas, com camadas desacopladas, portas explícitas e pipeline padronizado de transformação de dados para DTOs públicos. A API pública deve expor somente uma camada de serviço/fachada já tratada, sem acesso do consumidor ao client HTTP.

## Technical Context

**Language/Version**: Go 1.22+  
**Primary Dependencies**: `net/http`, `context`, `encoding/json`, `testing`, opcionalmente `httptest`  
**Storage**: N/A  
**Testing**: `go test` com testes unitários e integração leve com `httptest`  
**Target Platform**: Biblioteca Go para Linux/macOS/Windows  
**Project Type**: Go library (SDK)  
**Performance Goals**: P95 de transformação in-memory < 10ms para payload de referência  
**Constraints**: Sem dependência direta de infraestrutura nas camadas de domínio/aplicação; exposição mínima de campos  
**Scale/Scope**: Estrutura base + 1 fluxo de referência (integração/mapeamento/formatação)

## Constitution Check

- Passa no gate de simplicidade: apenas abstrações necessárias (ports e strategies com uso real).
- Passa no gate de desacoplamento: dependências unidirecionais (`domain <- application <- infra`).
- Passa no gate de testabilidade: use cases testáveis com doubles sem HTTP real.
- Passa no gate de rastreabilidade: tarefas mapeadas a histórias e requisitos.

## Project Structure

### Documentation (this feature)

```text
specs/001-project-scaffolding/
├── plan.md
├── spec.md
└── tasks.md
```

### Source Code (repository root)

```text
cmd/
└── examples/
    └── pokemon/
        └── main.go

internal/
├── domain/
│   ├── pokemon.go
│   └── errors.go
├── application/
│   └── pokemon/
│       ├── ports.go
│       ├── service.go
│       ├── mapper.go
│       └── formatter.go
└── infra/
    └── pokeapi/
        ├── client.go
        ├── dto.go
        └── adapter.go

pkg/
└── pokemon/
    ├── service.go
    └── models.go

docs/
└── architecture.md

tests/
└── application/
    └── pokemon/
        ├── service_test.go
        ├── mapper_test.go
        └── formatter_test.go

go.mod
README.md
```

**Structure Decision**: SDK Go única com separação entre API pública (`pkg/`), implementação interna (`internal/`) e documentação (`docs/`). `cmd/examples` fornece executável de referência sem acoplar à API pública.

## Boundary Definitions

- `internal/domain`: entidades e erros de negócio sem import de infra.
- `internal/application`: casos de uso, portas e estratégias de transformação/formatação; coordena fluxo.
- `internal/infra`: adaptação para APIs externas (HTTP, serialização, retry/policies).
- `pkg/pokemon`: camada pública de serviço/fachada estável, converte chamadas do usuário para use cases internos sem expor client HTTP.
- `docs/architecture.md`: contratos de dependência e regras para extensão.

## Traceability Matrix

- `FR-001`, `FR-002`, `FR-010`, `FR-011` -> Estrutura de diretórios + wiring de dependência em `pkg/pokemon/service.go`.
- `FR-003`, `FR-006` -> `internal/application/pokemon/ports.go`, `service.go` e tipos de erro.
- `FR-004`, `FR-009` -> `service.go` + `mapper.go` garantindo seleção de campos.
- `FR-005` -> `formatter.go` com estratégias de saída.
- `FR-007` -> `docs/architecture.md` + README.
- `FR-008` -> testes em `tests/application/pokemon/*.go`.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| Estratégia de formatação | Necessário para múltiplas visões da mesma entidade | Duplicar use cases geraria alto acoplamento e manutenção cara |
