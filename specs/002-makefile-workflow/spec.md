# Feature Specification: Makefile Workflow

**Feature Branch**: `[002-makefile-workflow]`  
**Created**: 2026-02-19  
**Status**: Draft  
**Input**: User description: "Preciso de uma nova feature, crie um makefile para rodar o projeto mais facilmente pelo terminal"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Execução rápida por comandos padronizados (Priority: P1)

Como desenvolvedor, quero usar comandos `make` curtos para tarefas comuns (testar e rodar exemplo), para reduzir tempo e erros ao trabalhar no terminal.

**Why this priority**: É o objetivo principal da feature e entrega valor imediato para uso diário.

**Independent Test**: Rodar `make help`, `make test` e `make run-example` com sucesso localmente.

**Acceptance Scenarios**:

1. **Given** o projeto clonado, **When** executo `make help`, **Then** vejo a lista de comandos disponíveis com breve descrição.
2. **Given** alterações no código, **When** executo `make test`, **Then** a suíte `go test ./...` roda de forma padronizada.
3. **Given** o exemplo da SDK, **When** executo `make run-example`, **Then** o executável de exemplo é iniciado sem precisar lembrar o comando completo.

---

### User Story 2 - Padronização de manutenção local (Priority: P2)

Como mantenedor, quero comandos de manutenção (`fmt`, `tidy`) no Makefile, para manter consistência do projeto.

**Why this priority**: Evita divergência entre membros do time e simplifica rotina de qualidade.

**Independent Test**: Rodar `make fmt` e `make tidy` e validar que os comandos Go são executados sem erro.

**Acceptance Scenarios**:

1. **Given** arquivos Go no repositório, **When** executo `make fmt`, **Then** `gofmt` é aplicado aos arquivos do projeto.
2. **Given** alterações de dependência, **When** executo `make tidy`, **Then** `go mod tidy` atualiza o módulo adequadamente.

---

### User Story 3 - Documentação de uso do Makefile (Priority: P3)

Como integrante novo do time, quero que o README explique o uso dos alvos do Makefile, para começar rapidamente sem buscar comandos internos.

**Why this priority**: Melhora onboarding e reduz dúvidas operacionais.

**Independent Test**: Seguir apenas README para executar os principais alvos com sucesso.

**Acceptance Scenarios**:

1. **Given** README atualizado, **When** um usuário consulta a seção de comandos, **Then** entende como usar os alvos principais.

---

### Edge Cases

- Como o `make fmt` deve se comportar quando não houver arquivos `.go`?
- Como o `make run-example` deve se comportar se o exemplo falhar por indisponibilidade de rede/API externa?
- Como evitar que alvos com nomes de arquivos reais sejam interpretados incorretamente (uso de `.PHONY`)?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: O sistema MUST incluir um `Makefile` na raiz do projeto.
- **FR-002**: O `Makefile` MUST definir um alvo `help` como guia de comandos disponíveis.
- **FR-003**: O `Makefile` MUST definir um alvo `test` para executar `go test ./...`.
- **FR-004**: O `Makefile` MUST definir um alvo `run-example` para executar `cmd/examples/pokemon/main.go`.
- **FR-005**: O `Makefile` MUST definir um alvo `fmt` para formatação de código Go.
- **FR-006**: O `Makefile` MUST definir um alvo `tidy` para manutenção de dependências do módulo.
- **FR-007**: O `Makefile` MUST usar `.PHONY` para os alvos operacionais.
- **FR-008**: O README MUST documentar os principais alvos e quando utilizá-los.

### Key Entities *(include if feature involves data)*

- **Make Target**: comando nomeado no `Makefile` que encapsula uma ação de desenvolvimento.
- **Developer Workflow Command**: operação de terminal padrão para build/teste/execução.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Usuário executa `make help`, `make test` e `make run-example` sem consultar comandos internos do projeto.
- **SC-002**: `make test` executa o mesmo comando base (`go test ./...`) usado manualmente.
- **SC-003**: README possui seção explícita de uso do Makefile com pelo menos os alvos `help`, `test`, `run-example`, `fmt` e `tidy`.
- **SC-004**: Alvos do Makefile são idempotentes para repetição local (sem efeitos colaterais inesperados).
