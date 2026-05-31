package pangolin

import (
	"errors"
	"net/http"
	"net/url"
)

type AuthTransport struct {
	transport http.RoundTripper
	apiKey    string
	baseURL   *url.URL
}

func NewAuthTransport(apiKey string, baseURL *url.URL) *AuthTransport {
	return &AuthTransport{
		transport: http.DefaultTransport,
		apiKey:    apiKey,
		baseURL:   baseURL,
	}
}

func (a *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Validate API key
	if a.apiKey == "" {
		return nil, errors.New("API key is required")
	}

	// Clone the request to avoid modifying the original, as required by the RoundTripper spec.
	req = req.Clone(req.Context())

	req.Header.Set("Authorization", "Bearer "+a.apiKey)

	req.URL.Scheme = a.baseURL.Scheme
	req.URL.Host = a.baseURL.Host

	return a.transport.RoundTrip(req)
}
