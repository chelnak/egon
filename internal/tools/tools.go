//Package tools contains a custom http client that retrieves
//tool data from a given url and returns it to the caller.
package tools

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	toolsURL = "https://puppetlabs.github.io/content-and-tooling-team/tools/list.json"
)

type ToolClient struct {
	client *http.Client
}

type Tool struct {
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}

func (m *ToolClient) GetTools(ctx context.Context) (*[]Tool, error) {
	req, err := http.NewRequest("GET", toolsURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := m.client.Do(req)

	if err != nil {
		return nil, err
	}

	response := new([]Tool)
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil
}

func NewToolClient(client *http.Client) *ToolClient {
	if client == nil {
		client = &http.Client{}
	}

	return &ToolClient{client: client}
}
