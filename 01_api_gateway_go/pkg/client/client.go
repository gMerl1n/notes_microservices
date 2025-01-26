package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type BaseClient struct {
	HTTPClient *http.Client
	Logger     *logging.Logger
}

func NewBaseClient(log *logging.Logger) *BaseClient {
	return &BaseClient{
		HTTPClient: &http.Client{},
		Logger:     log,
	}
}

func (c *BaseClient) SendRequest(req *http.Request) (*APIResponse, error) {
	if c.HTTPClient == nil {
		return nil, errors.New("no http client")
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request. error: %w", err)
	}

	apiResponse := APIResponse{
		IsOk:     true,
		response: response,
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		apiResponse.IsOk = false
		// if an error, read body so close it
		defer response.Body.Close()

		var apiErr APIError
		if err = json.NewDecoder(response.Body).Decode(&apiErr); err == nil {
			apiResponse.Error = apiErr
		}
	}

	return &apiResponse, nil
}

func (c *BaseClient) Close() error {
	c.HTTPClient = nil
	return nil
}
