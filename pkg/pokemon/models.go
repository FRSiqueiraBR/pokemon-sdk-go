package pokemon

// Format controls output projection for the same pokemon resource.
type Format string

const (
	FormatSummary  Format = "summary"
	FormatDetailed Format = "detailed"
)

// Pokemon is the stable public contract returned by the SDK.
type Pokemon struct {
	ID             int
	Name           string
	PrimaryType    string
	Types          []string
	Height         int
	Weight         int
	BaseExperience int
}
