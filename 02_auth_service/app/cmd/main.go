package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/iriskin77/notes_microservices/app/internal/auth"
	"github.com/iriskin77/notes_microservices/app/internal/config"
	"github.com/iriskin77/notes_microservices/app/pkg/db"
	"github.com/iriskin77/notes_microservices/app/pkg/jwt"
	"github.com/iriskin77/notes_microservices/app/pkg/logging"
	"github.com/iriskin77/notes_microservices/app/pkg/redis_client"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	// initializing config
	conf := config.LoadConfig()

	// initializing logger
	logger := logging.SetupLogger(conf.Env)

	// initializin db
	confDB := db.NewPostgresConfig(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("SSLMODE"),
	)

	ctx := context.Background()
	client, err := db.NewPostgresDB(ctx, confDB)

	if err != nil {
		fmt.Println(err.Error())
	}

	// initializing Redis

	i, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisConfig := redis_client.NewRedisConfig(
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		i,
	)

	clientRedis, err := redis_client.NewRedisClient(redisConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	storeRedis := auth.NewRedisStore(clientRedis)

	// initializing tokenManager to generate JWT

	// accessTokenTTL  int    `yaml:"access_tokenTTL"`
	// refreshTokenTTL int    `yaml:"refresh_tokenTTL"`

	tokenManager, err := jwt.NewManager(conf.JWTSecret, time.Duration(conf.AccessTokenTTL)*time.Minute, time.Duration(conf.RefreshTokenTTL)*time.Minute)
	if err != nil {
		fmt.Println(err.Error())
	}

	// initializing server
	repo := auth.NewRepository(client, logger)
	service := auth.NewService(repo, tokenManager, storeRedis, logger,
		time.Duration(conf.AccessTokenTTL)*time.Minute,
		time.Duration(conf.RefreshTokenTTL)*time.Minute)

	h := auth.NewHandler(service, tokenManager, logger)

	router := mux.NewRouter()

	h.RegisterHandlers(router)

	//srv, err := server.NewHttpServer(conf.Port)

	if err != nil {
		fmt.Println(err.Error())
	}

	logger.Info("starting application", slog.String("env", conf.Env), slog.Any("cfg", conf))

	srv, err := NewHttpServer(router, conf)

	if err != nil {
		fmt.Println(err.Error())
	}

	srv.ListenAndServe()

}

func NewHttpServer(router *mux.Router, cfg *config.Config) (*http.Server, error) {

	return &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}, nil
}
