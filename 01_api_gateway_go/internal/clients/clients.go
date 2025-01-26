package clients

import (
	"github.com/gMerl1n/notes_microservices/internal/clients/auth_client"
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type Client struct {
	UserClient auth_client.IClientUser
}

func NewClient(
	baseClient *client.BaseClient,
	log *logging.Logger,
	configAuthServer *config.ConfigAuthServer) *Client {

	return &Client{
		UserClient: auth_client.NewClientUser(baseClient, log, configAuthServer),
	}
}
