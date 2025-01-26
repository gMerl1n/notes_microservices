package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, log *logging.Logger, conf *config.Config, validator *validator.Validate) (*http.Server, error) {

	repo := repository.NewRepositoryUser(db, log, conf.User)
	serv := services.NewServiceUser(repo, tokenManager, redisUser, log, time.Duration(conf.Token.AccessTokenTTL), time.Duration(conf.Token.RefreshTokenTTL))
	h := handlers.NewHandlerUser(serv, tokenManager, log, validator)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/create_user", h.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/login_user", h.LoginUser).Methods("POST")
	router.HandleFunc("/api/v1/refresh_token", h.RefreshTokens).Methods("POST")

	return &http.Server{
		Addr:    conf.Server.Port,
		Handler: router,
	}, nil

}
