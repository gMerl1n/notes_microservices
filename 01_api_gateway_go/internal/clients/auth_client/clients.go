package auth_client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type IClientUser interface {
	CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (CreateUserResponse, error)
	LoginUser(ctx context.Context, email, password string) (TokensResponse, error)
	RefreshTokens(ctx context.Context, refreshToken string) (TokensResponse, error)
}

type ClientUser struct {
	baseClient *client.BaseClient
	authServer *config.ConfigAuthServer
	logger     *logging.Logger
}

func NewClientUser(
	baseClient *client.BaseClient,
	log *logging.Logger,
	configAuthServer *config.ConfigAuthServer) *ClientUser {

	return &ClientUser{
		baseClient: baseClient,
		authServer: configAuthServer,
		logger:     log,
	}
}

func (c *ClientUser) CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (CreateUserResponse, error) {

	var newUser CreateUserResponse

	data := url.Values{}
	data.Add("name", name)
	data.Add("surname", surname)
	data.Add("email", email)
	data.Add("password", password)
	data.Add("repeat_password", repeatPassword)
	data.Add("age", strconv.Itoa(age))

	req, err := http.NewRequest(http.MethodPost, c.authServer.UrlCreateUser, strings.NewReader(data.Encode()))
	if err != nil {
		return newUser, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return newUser, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		if err = json.NewDecoder(response.Body()).Decode(&newUser); err != nil {
			return newUser, fmt.Errorf("failed to decode body due to error %w", err)
		}
		return newUser, nil
	}
	return newUser, nil
}

func (c *ClientUser) LoginUser(ctx context.Context, email, password string) (TokensResponse, error) {

	var tokens TokensResponse

	data := url.Values{}
	data.Add("email", email)
	data.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, c.authServer.UrlCreateUser, strings.NewReader(data.Encode()))
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
		if err = json.NewDecoder(response.Body()).Decode(&tokens); err != nil {
			return tokens, fmt.Errorf("failed to decode body due to error %w", err)
		}
		return tokens, nil
	}
	return tokens, nil

}

func (c *ClientUser) RefreshTokens(ctx context.Context, refreshToken string) (TokensResponse, error) {

	var tokens TokensResponse

	data := url.Values{}
	data.Add("refresh_token", refreshToken)

	req, err := http.NewRequest(http.MethodPost, c.authServer.UrlCreateUser, strings.NewReader(data.Encode()))
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
		if err = json.NewDecoder(response.Body()).Decode(&tokens); err != nil {
			return tokens, fmt.Errorf("failed to decode body due to error %w", err)
		}
		return tokens, nil
	}
	return tokens, nil

}
