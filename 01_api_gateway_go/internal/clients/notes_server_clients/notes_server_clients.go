package notes_server_clients

import (
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type NotesServerClients struct {
	NotesClient      IClientNotes
	CategoriesClient IClientCategories
}

func NewNotesServerClients(baseClient *client.BaseClient, log *logging.Logger, config *config.Config) *NotesServerClients {
	return &NotesServerClients{
		NotesClient:      NewClientNotes(baseClient, log, config.NotesServer),
		CategoriesClient: NewClientCategories(baseClient, log, config.NotesServer),
	}
}
