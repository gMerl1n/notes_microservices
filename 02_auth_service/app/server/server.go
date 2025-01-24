package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gMerl1n/notes_microservices/app/internal/config"
	"github.com/gMerl1n/notes_microservices/app/internal/handlers"
	"github.com/gMerl1n/notes_microservices/app/internal/repository"
	"github.com/gMerl1n/notes_microservices/app/internal/services"
	"github.com/gMerl1n/notes_microservices/app/pkg/db"
	"github.com/gMerl1n/notes_microservices/app/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
	"github.com/gMerl1n/notes_microservices/app/pkg/redis_client"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, log *logging.Logger, conf *config.Config, tokenManager jwt.TokenManager, validator *validator.Validate) (*http.Server, error) {

	db, err := db.NewPostgresDB(ctx, conf.Postgres)

	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := redis_client.NewRedisClient(ctx, conf.Redis)
	if err != nil {
		log.Fatal("Failed to initialize Redis")
		return nil, err
	}

	redisUser := repository.NewRedisStoreUser(redisClient)

	repo := repository.NewRepositoryUser(db, log, conf.User)
	serv := services.NewServiceUser(repo, tokenManager, redisUser, log, time.Duration(conf.Token.AccessTokenTTL), time.Duration(conf.Token.RefreshTokenTTL))
	h := handlers.NewHandlerUser(serv, tokenManager, log, validator)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/create_user", h.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/login_user", h.LoginUser).Methods("POST")

	return &http.Server{
		Addr:    conf.Server.Port,
		Handler: router,
	}, nil

}
