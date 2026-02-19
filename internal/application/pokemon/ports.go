package pokemon

import "context"

// ExternalPokemonData is the application-facing representation from gateway adapters.
type ExternalPokemonData struct {
	ID             int
	Name           string
	Types          []string
	Height         int
	Weight         int
	BaseExperience int
}

// ExternalPokemonGateway defines the boundary for upstream API access.
type ExternalPokemonGateway interface {
	FetchPokemon(ctx context.Context, name string) (ExternalPokemonData, error)
}
