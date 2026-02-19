package pokemon

import (
	"fmt"
	"strings"

	"github.com/frsiqueira/pokemon-sdk-go/internal/domain"
)

func mapExternalToCore(data ExternalPokemonData) (domain.PokemonCore, error) {
	name := strings.TrimSpace(data.Name)
	if data.ID <= 0 || name == "" {
		return domain.PokemonCore{}, &domain.IntegrationError{
			Kind:      domain.ErrKindInvalidPayload,
			Operation: "map_external_to_core",
			Cause:     fmt.Errorf("missing required fields"),
		}
	}

	types := make([]string, 0, len(data.Types))
	for _, t := range data.Types {
		trimmed := strings.TrimSpace(t)
		if trimmed != "" {
			types = append(types, strings.ToLower(trimmed))
		}
	}

	return domain.PokemonCore{
		ID:             data.ID,
		Name:           strings.ToLower(name),
		Types:          types,
		Height:         data.Height,
		Weight:         data.Weight,
		BaseExperience: data.BaseExperience,
	}, nil
}
