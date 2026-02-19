package pokemon

import (
	"context"
	"fmt"
	"strings"

	"github.com/frsiqueira/pokemon-sdk-go/internal/domain"
)

// Service coordinates gateway access, mapping, and output formatting.
type Service struct {
	gateway    ExternalPokemonGateway
	formatters map[string]Formatter
}

func NewService(gateway ExternalPokemonGateway, formatters map[string]Formatter) *Service {
	if formatters == nil {
		formatters = defaultFormatters()
	}
	return &Service{gateway: gateway, formatters: formatters}
}

func (s *Service) GetPokemon(ctx context.Context, name string, format string) (PokemonView, error) {
	if strings.TrimSpace(name) == "" {
		return PokemonView{}, &domain.IntegrationError{
			Kind:      domain.ErrKindInvalidPayload,
			Operation: "get_pokemon",
			Cause:     fmt.Errorf("pokemon name is required"),
		}
	}

	data, err := s.gateway.FetchPokemon(ctx, name)
	if err != nil {
		return PokemonView{}, err
	}

	core, err := mapExternalToCore(data)
	if err != nil {
		return PokemonView{}, err
	}

	formatter := s.resolveFormatter(format)
	return formatter.Format(core), nil
}

func (s *Service) resolveFormatter(format string) Formatter {
	key := strings.ToLower(strings.TrimSpace(format))
	if key == "" {
		key = FormatSummary
	}
	formatter, ok := s.formatters[key]
	if !ok {
		return s.formatters[FormatSummary]
	}
	return formatter
}

func defaultFormatters() map[string]Formatter {
	return map[string]Formatter{
		FormatSummary:  SummaryFormatter{},
		FormatDetailed: DetailedFormatter{},
	}
}
