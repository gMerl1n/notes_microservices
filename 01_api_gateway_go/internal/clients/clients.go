package clients

import (
	"github.com/gMerl1n/notes_microservices/internal/clients/notes_server_clients"
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type Client struct {
	UserClient       IClientUser
	NotesClient      notes_server_clients.IClientNotes
	CategoriesClient notes_server_clients.IClientCategories
}

func NewClient(baseClient *client.BaseClient, log *logging.Logger, config *config.Config) *Client {

	return &Client{
		UserClient:       NewClientUser(baseClient, log, config.AuthServer),
		NotesClient:      notes_server_clients.NewClientNotes(baseClient, log, config.NotesServer),
		CategoriesClient: notes_server_clients.NewClientCategories(baseClient, log, config.NotesServer),
	}
}
