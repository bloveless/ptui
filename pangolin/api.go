package pangolin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type API struct {
	apiKey  string
	baseURL *url.URL
	client  http.Client
}

func NewAPI(apiKey string, baseURL *url.URL) API {
	client := http.Client{Timeout: 30 * time.Second, Transport: NewAuthTransport(apiKey, baseURL)}
	return API{
		client: client,
	}
}

func (a *API) ListOrgs() (ListOrgsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "/v1/orgs", http.NoBody)
	if err != nil {
		return ListOrgsResponse{}, fmt.Errorf("unable to generate list orgs request: %w", err)
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return ListOrgsResponse{}, fmt.Errorf("unable to execute request to list orgs: %w", err)
	}
	defer resp.Body.Close()
	listOrgsResponse := ListOrgsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&listOrgsResponse); err != nil {
		return ListOrgsResponse{}, fmt.Errorf("unable to decode response from list orgs endpoint: %w", err)
	}
	if listOrgsResponse.Error {
		return ListOrgsResponse{}, fmt.Errorf("detected error from list orgs response [status: %d]: %s", listOrgsResponse.Status, listOrgsResponse.Message)
	}
	return listOrgsResponse, nil
}
