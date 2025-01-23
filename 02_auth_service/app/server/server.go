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

	"github.com/gorilla/mux"
)

func NewHttpServer(
	ctx context.Context,
	log *logging.Logger,
	conf *config.Config,
	postgres *db.ConfigPostgres,
	redisConf *redis_client.ConfigRedis,
	BindAddr string,
	tokenManager jwt.TokenManager,
) (*http.Server, error) {

	// db, err := repository.NewRepositoryUser(ctx, postgres)
	// if err != nil {
	// 	log.Fatal("Failed to initialize DB")
	// 	return nil, err
	// }

	db, err := db.NewPostgresDB(ctx, postgres)

	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := redis_client.NewRedisClient(redisConf)
	if err != nil {
		log.Fatal("Failed to initialize Redis")
		return nil, err
	}

	redisUser := repository.NewRedisStoreUser(redisClient)

	repo := repository.NewRepositoryUser(db, log)
	serv := services.NewServiceUser(repo, tokenManager, redisUser, log, time.Duration(conf.AccessTokenTTL), time.Duration(conf.RefreshTokenTTL))
	h := handlers.NewHandlerUser(serv, tokenManager, log)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/create_user", h.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/login_user", h.LoginUser).Methods("POST")

	return &http.Server{
		Addr:    BindAddr,
		Handler: router,
	}, nil

}
