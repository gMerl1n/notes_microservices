package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type IClientUser interface {
	CreateUser(ctx context.Context, createUser *models.CreateUserRequest) ([]byte, error)
	LoginUser(ctx context.Context, loginUser *models.LoginUserRequest) ([]byte, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*models.TokensResponse, error)
}

type ClientUser struct {
	baseClient *client.BaseClient
	authServer *config.ConfigAuthServer
	logger     *logging.Logger
}

func NewClientUser(baseClient *client.BaseClient, log *logging.Logger, configAuthServer *config.ConfigAuthServer) *ClientUser {

	return &ClientUser{
		baseClient: baseClient,
		authServer: configAuthServer,
		logger:     log,
	}
}

func (c *ClientUser) CreateUser(ctx context.Context, createUser *models.CreateUserRequest) ([]byte, error) {

	dataByes, err := json.Marshal(createUser)
	if err != nil {
		c.logger.Warn("Failed to marshal new user data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.authServer.UrlCreateUser, bytes.NewBuffer(dataByes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		responseBytes, err := response.ReadBody()
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	} else {
		c.logger.Warn("Bad request", response.StatusCode())
		return nil, fmt.Errorf("bad request %d", response.StatusCode())
	}
}

func (c *ClientUser) LoginUser(ctx context.Context, loginUser *models.LoginUserRequest) ([]byte, error) {

	var tokens []byte

	dataBytes, err := json.Marshal(loginUser)
	if err != nil {
		return tokens, fmt.Errorf("failed to marshal dto")
	}

	req, err := http.NewRequest(http.MethodPost, c.authServer.UrlLoginUser, bytes.NewBuffer(dataBytes))
	if err != nil {
		return tokens, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return tokens, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		c.logger.Info("Good response", response.StatusCode())
		tokensBytes, err := response.ReadBody()
		if err != nil {
			return nil, err
		}
		return tokensBytes, nil
	} else {
		c.logger.Warn("Bad request", response.StatusCode())
		return nil, fmt.Errorf("bad request %d", response.StatusCode())
	}

}

func (c *ClientUser) RefreshTokens(ctx context.Context, refreshToken string) (*models.TokensResponse, error) {

	var tokens models.TokensResponse

	data := url.Values{}
	data.Add("refresh_token", refreshToken)

	req, err := http.NewRequest(http.MethodPost, c.authServer.UrlCreateUser, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		if err = json.NewDecoder(response.Body()).Decode(&tokens); err != nil {
			return &tokens, fmt.Errorf("failed to decode body due to error %w", err)
		}
		return nil, nil
	}
	return &tokens, nil

}
