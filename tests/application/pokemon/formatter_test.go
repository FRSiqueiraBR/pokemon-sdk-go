package pokemon_test

import (
	"context"
	"testing"

	app "github.com/frsiqueira/pokemon-sdk-go/internal/application/pokemon"
)

type formatterFakeGateway struct{}

func (formatterFakeGateway) FetchPokemon(_ context.Context, _ string) (app.ExternalPokemonData, error) {
	return app.ExternalPokemonData{
		ID:             6,
		Name:           "Charizard",
		Types:          []string{"fire", "flying"},
		Height:         17,
		Weight:         905,
		BaseExperience: 267,
	}, nil
}

func TestSummaryFormatReturnsReducedFields(t *testing.T) {
	svc := app.NewService(formatterFakeGateway{}, nil)

	result, err := svc.GetPokemon(context.Background(), "charizard", app.FormatSummary)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Height != 0 || result.Weight != 0 || result.BaseExperience != 0 {
		t.Fatalf("summary format must hide detailed fields: %+v", result)
	}
}

func TestDetailedFormatReturnsExpandedFields(t *testing.T) {
	svc := app.NewService(formatterFakeGateway{}, nil)

	result, err := svc.GetPokemon(context.Background(), "charizard", app.FormatDetailed)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Height == 0 || result.Weight == 0 || result.BaseExperience == 0 {
		t.Fatalf("detailed format should expose detailed fields: %+v", result)
	}
	if len(result.Types) != 2 {
		t.Fatalf("expected 2 types, got %+v", result.Types)
	}
}
