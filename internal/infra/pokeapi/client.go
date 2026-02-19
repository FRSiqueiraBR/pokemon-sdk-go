package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/frsiqueira/pokemon-sdk-go/internal/domain"
)

const defaultBaseURL = "https://pokeapi.co/api/v2"

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client fetches pokemon data from PokeAPI.
type Client struct {
	baseURL string
	http    doer
}

type ClientConfig struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(cfg ClientConfig) *Client {
	baseURL := strings.TrimRight(strings.TrimSpace(cfg.BaseURL), "/")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	var httpClient doer
	if cfg.HTTPClient != nil {
		httpClient = cfg.HTTPClient
	} else {
		httpClient = &http.Client{Timeout: 5 * time.Second}
	}

	return &Client{baseURL: baseURL, http: httpClient}
}

func (c *Client) FetchPokemonRaw(ctx context.Context, name string) (pokemonResponse, error) {
	endpoint := fmt.Sprintf("%s/pokemon/%s", c.baseURL, url.PathEscape(strings.ToLower(strings.TrimSpace(name))))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return pokemonResponse{}, &domain.IntegrationError{Kind: domain.ErrKindInvalidPayload, Operation: "build_request", Cause: err}
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return pokemonResponse{}, &domain.IntegrationError{Kind: domain.ErrKindTimeout, Operation: "http_do", Cause: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return pokemonResponse{}, &domain.IntegrationError{
			Kind:       domain.ErrKindUpstreamStatus,
			Operation:  "read_response",
			StatusCode: resp.StatusCode,
		}
	}

	var payload pokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return pokemonResponse{}, &domain.IntegrationError{Kind: domain.ErrKindDecode, Operation: "decode_response", Cause: err}
	}

	return payload, nil
}
