---

description: "Task list for project scaffolding feature"

---

# Tasks: Project Scaffolding

**Input**: Design documents from `/specs/001-project-scaffolding/`
**Prerequisites**: plan.md, spec.md
**Tests**: Incluídos por exigência explícita do escopo.
**Organization**: Tasks agrupadas por user story (US1, US2, US3).

## Format: `[ID] [P?] [Story] Description`

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Inicialização do módulo Go e estrutura base de pastas.

- [ ] T001 [US1] Criar `go.mod` com módulo do projeto na raiz.
- [ ] T002 [P] [US1] Criar estrutura de diretórios definida no plano (`cmd/`, `internal/`, `pkg/`, `docs/`, `tests/`).
- [ ] T003 [P] [US1] Criar arquivos iniciais de package (`doc.go` quando necessário) para tornar limites explícitos.

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Definir contratos e fundações que bloqueiam todas as histórias.

- [ ] T004 [US1] Definir entidades de domínio em `internal/domain/pokemon.go`.
- [ ] T005 [P] [US1] Definir erros padronizados em `internal/domain/errors.go`.
- [ ] T006 [US1] Definir portas de integração em `internal/application/pokemon/ports.go`.
- [ ] T007 [US1] Criar camada pública de serviço/fachada em `pkg/pokemon/service.go` com injeção de dependências.
- [ ] T008 [P] [US1] Criar modelo público em `pkg/pokemon/models.go` sem vazar payload externo.

**Checkpoint**: Fundação pronta para implementar fluxo de integração e transformação.

---

## Phase 3: User Story 1 - Estrutura base e limites arquiteturais (Priority: P1) 🎯 MVP

**Goal**: Entregar baseline compilável com fronteiras claras e sem acoplamento indevido.

**Independent Test**: `go test ./...` passa com smoke tests e verificação de imports.

### Tests for User Story 1

- [ ] T009 [P] [US1] Criar teste de compilação/wiring em `tests/application/pokemon/service_test.go` com fake gateway.
- [ ] T010 [US1] Validar que `internal/application` não importa `internal/infra` (cheque automatizado simples via script/teste).

### Implementation for User Story 1

- [ ] T011 [US1] Implementar `internal/application/pokemon/service.go` coordenando portas e domínio.
- [ ] T012 [US1] Criar adaptador mínimo de exemplo em `internal/infra/pokeapi/adapter.go`.

**Checkpoint**: Arquitetura base validada e pronta para fluxo funcional.

---

## Phase 4: User Story 2 - Pipeline de integração e transformação de dados (Priority: P2)

**Goal**: Implementar fluxo padrão de fetch, mapeamento, filtro e formatação.

**Independent Test**: Com payload mockado, serviço retorna DTO público filtrado.

### Tests for User Story 2

- [ ] T013 [P] [US2] Criar testes de mapeamento em `tests/application/pokemon/mapper_test.go`.
- [ ] T014 [P] [US2] Criar testes de formatação em `tests/application/pokemon/formatter_test.go`.
- [ ] T015 [US2] Criar teste de serviço para descarte de campos extras em `tests/application/pokemon/service_test.go`.
- [ ] T015A [US2] Criar teste garantindo que o consumidor usa apenas a fachada pública e não acessa client HTTP diretamente.

### Implementation for User Story 2

- [ ] T016 [US2] Implementar mapeador em `internal/application/pokemon/mapper.go`.
- [ ] T017 [US2] Implementar estratégias de formatação em `internal/application/pokemon/formatter.go`.
- [ ] T018 [US2] Implementar cliente HTTP e DTO de infra em `internal/infra/pokeapi/client.go` e `internal/infra/pokeapi/dto.go`.
- [ ] T019 [US2] Integrar pipeline completo em `internal/application/pokemon/service.go`.

**Checkpoint**: Fluxo funcional completo e testável sem API real (via doubles).

---

## Phase 5: User Story 3 - Documentação operacional da arquitetura (Priority: P3)

**Goal**: Documentar decisões, limites e guia de extensão.

**Independent Test**: Desenvolvedor novo consegue adicionar integração exemplo seguindo docs.

### Implementation for User Story 3

- [ ] T020 [US3] Documentar arquitetura e regras de dependência em `docs/architecture.md`.
- [ ] T021 [US3] Atualizar `README.md` com visão do scaffold, execução de testes e exemplo de uso.
- [ ] T022 [US3] Documentar checklist de nova integração (porta, adaptador, mapper, testes) em `docs/architecture.md`.

**Checkpoint**: Documentação alinhada ao código e pronta para onboarding.

---

## Phase N: Polish & Cross-Cutting Concerns

- [ ] T023 [P] Revisar nomenclaturas e consistência entre camadas.
- [ ] T024 Executar `go test ./...` e corrigir regressões.
- [ ] T025 Garantir rastreabilidade final FR -> arquivos -> testes no fechamento da feature.

---

## Dependencies & Execution Order

- Setup (Phase 1) -> Foundational (Phase 2) -> US1 -> US2 -> US3 -> Polish.
- US2 depende de contratos definidos em US1.
- US3 depende da estrutura e fluxo implementados em US1/US2 para documentação fiel.

## Parallel Opportunities

- T002, T003 em paralelo.
- T005 e T008 em paralelo.
- T013 e T014 em paralelo.
- T023 pode ocorrer em paralelo com correções finais.

## Implementation Strategy

1. Entregar MVP arquitetural (US1) sem integração externa completa.
2. Evoluir pipeline funcional com mapeamento/formatação (US2).
3. Consolidar documentação de manutenção e extensão (US3).
4. Fechar com testes e rastreabilidade.
