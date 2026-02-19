package pokemon

import "github.com/frsiqueira/pokemon-sdk-go/internal/domain"

const (
	FormatSummary  = "summary"
	FormatDetailed = "detailed"
)

// PokemonView is the application output contract consumed by public API layer.
type PokemonView struct {
	ID             int
	Name           string
	PrimaryType    string
	Types          []string
	Height         int
	Weight         int
	BaseExperience int
}

// Formatter defines output projections for the same normalized entity.
type Formatter interface {
	Format(p domain.PokemonCore) PokemonView
}

type SummaryFormatter struct{}

type DetailedFormatter struct{}

func (SummaryFormatter) Format(p domain.PokemonCore) PokemonView {
	primaryType := ""
	if len(p.Types) > 0 {
		primaryType = p.Types[0]
	}

	return PokemonView{
		ID:          p.ID,
		Name:        p.Name,
		PrimaryType: primaryType,
	}
}

func (DetailedFormatter) Format(p domain.PokemonCore) PokemonView {
	typesCopy := make([]string, len(p.Types))
	copy(typesCopy, p.Types)

	return PokemonView{
		ID:             p.ID,
		Name:           p.Name,
		PrimaryType:    firstType(p.Types),
		Types:          typesCopy,
		Height:         p.Height,
		Weight:         p.Weight,
		BaseExperience: p.BaseExperience,
	}
}

func firstType(types []string) string {
	if len(types) == 0 {
		return ""
	}
	return types[0]
}
