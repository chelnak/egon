//Package modules contains a custom http client that retrieves
//module data from a given URL and returns it to the caller.
package modules

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	modulesURL = "https://puppetlabs.github.io/content-and-tooling-team/modules/list.json"
)

type ModuleClient struct {
	client *http.Client
}

type Module struct {
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}

func (m *ModuleClient) GetSupportedModules(ctx context.Context) (*[]Module, error) {
	req, err := http.NewRequest("GET", modulesURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := m.client.Do(req)

	if err != nil {
		return nil, err
	}

	response := new([]Module)
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil
}

func NewModuleClient(client *http.Client) *ModuleClient {
	if client == nil {
		client = &http.Client{}
	}

	return &ModuleClient{client: client}
}
