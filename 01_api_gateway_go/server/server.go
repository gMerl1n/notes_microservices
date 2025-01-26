package server

import (
	"context"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/clients"
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/handlers"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, log *logging.Logger, conf *config.Config, validator *validator.Validate, jwtParser jwt.ITokenParser) (*http.Server, error) {

	baseClient := client.NewBaseClient(log)

	// Инициализация клиентов
	clients := clients.NewClient(baseClient, log, conf.AuthServer)

	// Инициализация ручек
	handlers := handlers.NewHandlers(clients.UserClient, jwtParser, validator, log)

	router := mux.NewRouter()

	// auth handlers
	router.HandleFunc("/api_gateway/v1/create_user", handlers.HandlersUser.CreateUser).Methods("POST")
	router.HandleFunc("/api_gateway/v1/login_user", handlers.HandlersUser.LoginUser).Methods("POST")
	router.HandleFunc("/api_gateway/v1/refresh_token", handlers.HandlersUser.RefreshTokens).Methods("POST")

	// notices handlers

	return &http.Server{
		Addr:    conf.Server.Port,
		Handler: router,
	}, nil

}
