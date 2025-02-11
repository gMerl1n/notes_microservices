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

type IClientNotes interface {
	CreateNote(ctx context.Context, userCreate *models.NoteCreateRequest) ([]byte, error)
	GetNoteByID(ctx context.Context, noteGetRequest *models.NoteGetRequestByID) ([]byte, error)
	GetNotes(ctx context.Context, userID int) ([]byte, error)
	RemoveNoteByID(ctx context.Context, noteRemoveRequest *models.NoteRemoveRequestByID) ([]byte, error)
	RemoveNotes(ctx context.Context, userID int) ([]byte, error)
}

type ClientNotes struct {
	baseClient  *client.BaseClient
	logger      *logging.Logger
	notesServer *config.ConfigNotesServer
}

func NewClientNotes(baseClient *client.BaseClient, log *logging.Logger, configNotesServer *config.ConfigNotesServer) *ClientNotes {
	return &ClientNotes{
		baseClient:  baseClient,
		logger:      log,
		notesServer: configNotesServer,
	}
}

func (c *ClientNotes) CreateNote(ctx context.Context, userCreate *models.NoteCreateRequest) ([]byte, error) {

	dataByes, err := json.Marshal(userCreate)
	if err != nil {
		c.logger.Warn("Failed to marshal new note data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlCreateNote, bytes.NewBuffer(dataByes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request note creation due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.baseClient.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request note creation due to error: %w", err)
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

func (c *ClientNotes) GetNoteByID(ctx context.Context, noteGetRequest *models.NoteGetRequestByID) ([]byte, error) {

	dataByes, err := json.Marshal(noteGetRequest)
	if err != nil {
		c.logger.Warn("Failed to marshal new user data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlGetNoteByID, bytes.NewBuffer(dataByes))
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

func (c *ClientNotes) GetNotes(ctx context.Context, userID int) ([]byte, error) {

	dataByes, err := json.Marshal(userID)
	if err != nil {
		c.logger.Warn("Failed to marshal new user data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlGetNotes, bytes.NewBuffer(dataByes))
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

func (c *ClientNotes) RemoveNoteByID(ctx context.Context, noteRemoveRequest *models.NoteRemoveRequestByID) ([]byte, error) {

	dataByes, err := json.Marshal(noteRemoveRequest)
	if err != nil {
		c.logger.Warn("Failed to marshal new user data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlRemoveNoteByID, bytes.NewBuffer(dataByes))
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

func (c *ClientNotes) RemoveNotes(ctx context.Context, userID int) ([]byte, error) {

	dataByes, err := json.Marshal(userID)
	if err != nil {
		c.logger.Warn("Failed to marshal new user data %w", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.notesServer.UrlRemoveNoteByID, bytes.NewBuffer(dataByes))
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
