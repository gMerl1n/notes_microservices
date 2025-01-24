package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/handlers"
	"github.com/gMerl1n/notes_microservices/internal/repository"
	"github.com/gMerl1n/notes_microservices/internal/services"
	"github.com/gMerl1n/notes_microservices/pkg/db"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/gMerl1n/notes_microservices/pkg/redis_client"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, log *logging.Logger, conf *config.Config, tokenManager jwt.TokenManager, validator *validator.Validate) (*http.Server, error) {

	db, err := db.NewPostgresDB(ctx, conf.Postgres)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf.Redis)
	redisClient, err := redis_client.NewRedisClient(ctx, conf.Redis)
	if err != nil {
		log.Fatal("Failed to initialize Redis", err)
		return nil, err
	}

	redisUser := repository.NewRedisStoreUser(redisClient, log)

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
