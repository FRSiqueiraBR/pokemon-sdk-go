# Architecture

This project follows a boundary-first Go architecture to keep integrations simple and low coupled.

## Layers

- `internal/domain`: core entities and domain-level integration errors.
- `internal/application`: use cases, gateway ports, mapping, and output format strategies.
- `internal/infra`: external API adapters and HTTP clients.
- `pkg/pokemon`: public service facade that returns already-processed public models.

## Public API Contract

- Consumers must use `pkg/pokemon.Service`.
- Consumers must not access HTTP client implementations directly.
- External payloads are converted and filtered before crossing the public boundary.

## Dependency Rule

- Allowed: `domain <- application <- infra` and `pkg -> application + infra (wiring only)`.
- Forbidden: `domain` importing `application/infra`, and `application` importing `infra`.

## Standard Flow

1. Fetch raw data from external API through gateway port.
2. Map raw data to normalized domain entity.
3. Apply output formatting strategy (`summary` or `detailed`).
4. Return only public DTO fields from `pkg/pokemon`.

## Error Handling

Integration failures are wrapped in `domain.IntegrationError` with explicit kind:

- `timeout`
- `upstream_status`
- `decode`
- `invalid_payload`

## New Integration Checklist

- Add a gateway port contract in `internal/application/<feature>/ports.go`.
- Implement adapter in `internal/infra/<provider>/`.
- Keep provider DTOs private to infra package.
- Map provider payload to normalized domain entity.
- Add formatter/view logic in application layer when needed.
- Expose only stable public models and service operations in `pkg/`.
- Add/adjust unit tests under `tests/application/`.
