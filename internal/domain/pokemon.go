package domain

// PokemonCore represents normalized business data independent from external APIs.
type PokemonCore struct {
	ID             int
	Name           string
	Types          []string
	Height         int
	Weight         int
	BaseExperience int
}
