package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type IClientUser interface {
	CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (*models.CreateUserResponse, error)
	LoginUser(ctx context.Context, loginUser *models.LoginUserRequest) ([]byte, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*models.TokensResponse, error)
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

func (c *ClientUser) CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (*models.CreateUserResponse, error) {

	var newUser models.CreateUserResponse

	data := url.Values{}
	data.Add("name", name)
	data.Add("surname", surname)
	data.Add("email", email)
	data.Add("password", password)
	data.Add("repeat_password", repeatPassword)
	data.Add("age", strconv.Itoa(age))

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
		if err = json.NewDecoder(response.Body()).Decode(&newUser); err != nil {
			return nil, fmt.Errorf("failed to decode body due to error %w", err)
		}
		return &newUser, nil
	}
	return &newUser, nil
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
	}

	c.logger.Warn("Bad request", response.StatusCode())

	return nil, err

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
