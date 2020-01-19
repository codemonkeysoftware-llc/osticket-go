package osticket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// APIClient represents a client of an osTicket API
type APIClient struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewAPIClient returns a new client for making API requests. baseURL should include
// the url to your installation of osTicket. Do not include the `/api/` portion.
func NewAPIClient(httpClient *http.Client, baseURL, apiKey string) *APIClient {
	return &APIClient{
		client:  httpClient,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// CreateTicket makes a ticket request to the osTicket API
func (api *APIClient) CreateTicket(cmd *CreateTicketCommand) error {
	url := fmt.Sprintf("%s/api/http.php/tickets.json", api.baseURL)

	marshaled, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(marshaled))
	if err != nil {
		return err
	}
	r.Header.Set("X-API-Key", api.apiKey)
	r.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}
	return nil
}
