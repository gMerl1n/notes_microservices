package notes_server_clients

import (
	"context"

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
	RemoveNotes(ctx context.Context, userID int)
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
	return nil, nil
}

func (c *ClientNotes) GetNoteByID(ctx context.Context, noteGetRequest *models.NoteGetRequestByID) ([]byte, error) {
	return nil, nil
}

func (c *ClientNotes) GetNotes(ctx context.Context, userID int) ([]byte, error) {
	return nil, nil
}

func (c *ClientNotes) RemoveNoteByID(ctx context.Context, noteRemoveRequest *models.NoteRemoveRequestByID) ([]byte, error) {
	return nil, nil
}

func (c *ClientNotes) RemoveNotes(ctx context.Context, userID int) {}
