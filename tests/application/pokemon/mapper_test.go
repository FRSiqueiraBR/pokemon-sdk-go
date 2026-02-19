package pokemon_test

import (
	"context"
	"testing"

	app "github.com/frsiqueira/pokemon-sdk-go/internal/application/pokemon"
)

type mapperFakeGateway struct {
	data app.ExternalPokemonData
	err  error
}

func (g mapperFakeGateway) FetchPokemon(_ context.Context, _ string) (app.ExternalPokemonData, error) {
	return g.data, g.err
}

func TestServiceMapsAndFiltersData(t *testing.T) {
	svc := app.NewService(mapperFakeGateway{data: app.ExternalPokemonData{
		ID:             25,
		Name:           "Pikachu",
		Types:          []string{"electric", ""},
		Height:         4,
		Weight:         60,
		BaseExperience: 112,
	}}, nil)

	result, err := svc.GetPokemon(context.Background(), "pikachu", app.FormatSummary)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.ID != 25 || result.Name != "pikachu" {
		t.Fatalf("unexpected base mapping: %+v", result)
	}
	if result.PrimaryType != "electric" {
		t.Fatalf("expected primary type electric, got %q", result.PrimaryType)
	}
	if len(result.Types) != 0 {
		t.Fatalf("summary should not expose full type list, got %+v", result.Types)
	}
}
