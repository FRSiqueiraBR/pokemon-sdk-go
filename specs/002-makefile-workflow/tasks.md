---

description: "Task list for Makefile workflow feature"

---

# Tasks: Makefile Workflow

**Input**: Design documents from `/specs/002-makefile-workflow/`
**Prerequisites**: plan.md, spec.md
**Tests**: Validação por execução direta dos alvos make e `go test ./...`.
**Organization**: Tasks agrupadas por user story (US1, US2, US3).

## Format: `[ID] [P?] [Story] Description`

## Phase 1: Setup (Shared Infrastructure)

- [ ] T001 [US1] Criar arquivo `Makefile` na raiz com cabeçalho e padrões iniciais.

---

## Phase 2: Foundational (Blocking Prerequisites)

- [ ] T002 [US1] Definir `.PHONY` para alvos operacionais (`help`, `test`, `run-example`, `fmt`, `tidy`).
- [ ] T003 [US1] Implementar alvo `help` com descrições dos comandos.

**Checkpoint**: Base do fluxo make pronta.

---

## Phase 3: User Story 1 - Execução rápida por comandos padronizados (Priority: P1) 🎯 MVP

**Goal**: Entregar comandos make para teste e execução de exemplo.

**Independent Test**: Rodar `make help`, `make test`, `make run-example`.

### Implementation for User Story 1

- [ ] T004 [US1] Implementar alvo `test` com `go test ./...` em `Makefile`.
- [ ] T005 [US1] Implementar alvo `run-example` para `go run ./cmd/examples/pokemon` em `Makefile`.
- [ ] T006 [US1] Validar os alvos no terminal local.

---

## Phase 4: User Story 2 - Padronização de manutenção local (Priority: P2)

**Goal**: Entregar comandos de manutenção de código e módulo.

**Independent Test**: Rodar `make fmt` e `make tidy`.

### Implementation for User Story 2

- [ ] T007 [US2] Implementar alvo `fmt` em `Makefile`.
- [ ] T008 [US2] Implementar alvo `tidy` em `Makefile`.
- [ ] T009 [US2] Validar repetibilidade dos alvos sem efeitos colaterais.

---

## Phase 5: User Story 3 - Documentação de uso do Makefile (Priority: P3)

**Goal**: Documentar uso rápido dos comandos make.

**Independent Test**: Usuário executa os alvos apenas com README.

### Implementation for User Story 3

- [ ] T010 [US3] Adicionar seção de comandos Make no `README.md`.
- [ ] T011 [US3] Incluir exemplos curtos de uso para `help`, `test`, `run-example`, `fmt`, `tidy`.

---

## Phase N: Polish & Cross-Cutting Concerns

- [ ] T012 [P] Revisar nomenclatura e mensagens de ajuda no `Makefile`.
- [ ] T013 Executar validação final: `make help`, `make test`, `make run-example`, `make fmt`, `make tidy`.
- [ ] T014 Atualizar rastreabilidade FR -> arquivos -> validação.

## Dependencies & Execution Order

- Phase 1 -> Phase 2 -> US1 -> US2 -> US3 -> Polish.
- US2 depende do `Makefile` base criado em US1.
- US3 depende dos alvos finais para documentação fiel.

## Parallel Opportunities

- T007 e T008 podem ser implementadas em paralelo.
- T010 e T011 podem ocorrer em paralelo após conclusão dos alvos.

## Implementation Strategy

1. Entregar MVP operacional (`help`, `test`, `run-example`).
2. Adicionar manutenção (`fmt`, `tidy`).
3. Finalizar com documentação e validação ponta a ponta.
