package pokemon

import (
	"context"
	"net/http"

	app "github.com/frsiqueira/pokemon-sdk-go/internal/application/pokemon"
	"github.com/frsiqueira/pokemon-sdk-go/internal/infra/pokeapi"
)

// Config wires SDK dependencies while keeping sensible defaults.
type Config struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Service is the only public SDK entrypoint and returns already-processed models.
type Service struct {
	appService *app.Service
}

// NewService creates the public facade that hides HTTP client details from SDK consumers.
func NewService(cfg Config) *Service {
	apiClient := pokeapi.NewClient(pokeapi.ClientConfig{
		BaseURL:    cfg.BaseURL,
		HTTPClient: cfg.HTTPClient,
	})
	gateway := pokeapi.NewGateway(apiClient)
	service := app.NewService(gateway, nil)

	return &Service{appService: service}
}

// GetPokemon fetches external data and returns a normalized, filtered public object.
func (s *Service) GetPokemon(ctx context.Context, name string, format Format) (Pokemon, error) {
	view, err := s.appService.GetPokemon(ctx, name, string(format))
	if err != nil {
		return Pokemon{}, err
	}

	return Pokemon{
		ID:             view.ID,
		Name:           view.Name,
		PrimaryType:    view.PrimaryType,
		Types:          view.Types,
		Height:         view.Height,
		Weight:         view.Weight,
		BaseExperience: view.BaseExperience,
	}, nil
}
