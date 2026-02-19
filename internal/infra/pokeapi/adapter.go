package pokeapi

import (
	"context"
	"sort"

	app "github.com/frsiqueira/pokemon-sdk-go/internal/application/pokemon"
)

// Gateway adapts PokeAPI client into the application gateway port.
type Gateway struct {
	client *Client
}

func NewGateway(client *Client) *Gateway {
	return &Gateway{client: client}
}

func (g *Gateway) FetchPokemon(ctx context.Context, name string) (app.ExternalPokemonData, error) {
	payload, err := g.client.FetchPokemonRaw(ctx, name)
	if err != nil {
		return app.ExternalPokemonData{}, err
	}

	sort.Slice(payload.Types, func(i, j int) bool {
		return payload.Types[i].Slot < payload.Types[j].Slot
	})

	types := make([]string, 0, len(payload.Types))
	for _, t := range payload.Types {
		types = append(types, t.Type.Name)
	}

	return app.ExternalPokemonData{
		ID:             payload.ID,
		Name:           payload.Name,
		Types:          types,
		Height:         payload.Height,
		Weight:         payload.Weight,
		BaseExperience: payload.BaseExperience,
	}, nil
}
