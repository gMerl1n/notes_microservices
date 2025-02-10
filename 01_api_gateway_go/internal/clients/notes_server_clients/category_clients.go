package notes_server_clients

import (
	"context"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type IClientCategories interface {
	CreateCategory(ctx context.Context, categoryCreate *models.CategoryCreateRequest) ([]byte, error)
	GetCategoryByID(ctx context.Context, categoryGetRequest *models.CategoryGetRequestByID) ([]byte, error)
	GetCategories(ctx context.Context, userID int) ([]byte, error)
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
	return nil, nil
}

func (c *ClientCategories) GetCategoryByID(ctx context.Context, categoryGetRequest *models.CategoryGetRequestByID) ([]byte, error) {
	return nil, nil
}

func (c *ClientCategories) GetCategories(ctx context.Context, userID int) ([]byte, error) {
	return nil, nil
}

func (c *ClientCategories) RemoveCategoryByID(ctx context.Context, categoryRemoveRequest *models.CategoryRemoveRequestByID) ([]byte, error) {
	return nil, nil
}
