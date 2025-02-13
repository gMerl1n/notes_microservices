package notes_server_clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type IClientCategories interface {
	CreateCategory(ctx context.Context, categoryCreate *models.CategoryCreateRequest) ([]byte, error)
	GetCategoryByID(ctx context.Context, categoryGetRequest *models.CategoryGetRequestByID) ([]byte, error)
	GetCategories(ctx context.Context, categoriesGetRequest *models.CategoriesGetRequest) ([]byte, error)
	RemoveCategoryByID(ctx context.Context, categoryRemoveRequest *models.CategoryRemoveRequestByID) ([]byte, error)
}

type ClientCategories struct {
	baseClient  *client.BaseClient
	logger      *logging.Logger
	notesServer *config.ConfigNotesServer
}

func NewClientCategories(baseClient *client.BaseClient, log *logging.Logger, notesServer *config.ConfigNotesServer) *ClientCategories {
	return &ClientCategories{
		baseClient:  baseClient,
		logger:      log,
		notesServer: notesServer,
	}
}

func (c *ClientCategories) CreateCategory(ctx context.Context, categoryCreate *models.CategoryCreateRequest) ([]byte, error) {

	dataByes, err := json.Marshal(categoryCreate)
	if err != nil {
		c.logger.Warn("failed to marshal new get note data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlCreateCategory, bytes.NewBuffer(dataByes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request get note by id due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request get note by due to error: %w", err)
	}

	if response.IsOk {
		responseBytes, err := response.ReadBody()
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	} else {
		c.logger.Warn("bad request", response.StatusCode())
		return nil, fmt.Errorf("bad request %d", response.StatusCode())
	}

}

func (c *ClientCategories) GetCategoryByID(ctx context.Context, categoryGetRequestByID *models.CategoryGetRequestByID) ([]byte, error) {

	dataByes, err := json.Marshal(categoryGetRequestByID)
	if err != nil {
		c.logger.Warn("failed to marshal new get note data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlGetCategoryByID, bytes.NewBuffer(dataByes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request get note by id due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request get note by due to error: %w", err)
	}

	if response.IsOk {
		responseBytes, err := response.ReadBody()
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	} else {
		c.logger.Warn("bad request", response.StatusCode())
		return nil, fmt.Errorf("bad request %d", response.StatusCode())
	}

}

func (c *ClientCategories) GetCategories(ctx context.Context, categoriesGetRequest *models.CategoriesGetRequest) ([]byte, error) {

	dataByes, err := json.Marshal(categoriesGetRequest)
	if err != nil {
		c.logger.Warn("failed to marshal new get notes data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlGetCategories, bytes.NewBuffer(dataByes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request get nots due to error: %w", err)
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
		c.logger.Warn("bad request", response.StatusCode())
		return nil, fmt.Errorf("bad request %d", response.StatusCode())
	}

}

func (c *ClientCategories) RemoveCategoryByID(ctx context.Context, categoryRemoveRequest *models.CategoryRemoveRequestByID) ([]byte, error) {

	dataByes, err := json.Marshal(categoryRemoveRequest)
	if err != nil {
		c.logger.Warn("Failed to marshal new remove notes data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, c.notesServer.UrlRemoveCategoryByID, bytes.NewBuffer(dataByes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request remove notes due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request remove notes due to error: %w", err)
	}

	if response.IsOk {
		responseBytes, err := response.ReadBody()
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	} else {
		c.logger.Warn("bad request", response.StatusCode())
		return nil, fmt.Errorf("bad request %d", response.StatusCode())
	}

}
