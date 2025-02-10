package clients

import (
	"github.com/gMerl1n/notes_microservices/internal/clients/notes_server_clients"
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type Client struct {
	UserClient        IClientUser
	NotesServerClient *notes_server_clients.NotesServerClients
}

func NewClient(baseClient *client.BaseClient, log *logging.Logger, config *config.Config) *Client {

	return &Client{
		UserClient:        NewClientUser(baseClient, log, config.AuthServer),
		NotesServerClient: notes_server_clients.NewNotesServerClients(baseClient, log, config),
	}
}
