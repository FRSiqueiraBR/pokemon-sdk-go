package pokemon_test

import (
	"context"
	"errors"
	"testing"

	app "github.com/frsiqueira/pokemon-sdk-go/internal/application/pokemon"
	"github.com/frsiqueira/pokemon-sdk-go/internal/domain"
)

type serviceFakeGateway struct {
	data app.ExternalPokemonData
	err  error
}

func (g serviceFakeGateway) FetchPokemon(_ context.Context, _ string) (app.ExternalPokemonData, error) {
	if g.err != nil {
		return app.ExternalPokemonData{}, g.err
	}
	return g.data, nil
}

func TestServicePropagatesGatewayError(t *testing.T) {
	inErr := &domain.IntegrationError{Kind: domain.ErrKindTimeout, Operation: "http_do"}
	svc := app.NewService(serviceFakeGateway{err: inErr}, nil)

	_, err := svc.GetPokemon(context.Background(), "pikachu", app.FormatSummary)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, inErr) {
		t.Fatalf("expected wrapped gateway error, got %v", err)
	}
}

func TestServiceFallsBackToSummaryWhenUnknownFormat(t *testing.T) {
	svc := app.NewService(serviceFakeGateway{data: app.ExternalPokemonData{
		ID:   1,
		Name: "bulbasaur",
		Types: []string{
			"grass",
			"poison",
		},
		Height:         7,
		Weight:         69,
		BaseExperience: 64,
	}}, nil)

	result, err := svc.GetPokemon(context.Background(), "bulbasaur", "unknown")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Height != 0 || result.Weight != 0 {
		t.Fatalf("unknown format should fallback to summary: %+v", result)
	}
}
