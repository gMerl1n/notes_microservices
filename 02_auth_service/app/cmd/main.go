package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gMerl1n/notes_microservices/app/internal/config"
	"github.com/gMerl1n/notes_microservices/app/pkg/db"
	"github.com/gMerl1n/notes_microservices/app/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
	"github.com/gMerl1n/notes_microservices/app/pkg/redis_client"
	"github.com/gMerl1n/notes_microservices/app/server"
	"github.com/joho/godotenv"
)

func main() {

	// initializing config
	conf := config.LoadConfig("./app/config/config.yml")

	// initializing logger
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	// load .env variables

	if err := godotenv.Load(); err != nil {
		logger.Fatal(err)
	}

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
	// client, err := db.NewPostgresDB(ctx, confDB)

	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// initializing Redis

	redisConfig, err := redis_client.NewRedisConfig(
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_DB"),
	)
	if err != nil {
		logger.Fatal(err)
	}

	// clientRedis, err := redis_client.NewRedisClient(redisConfig)
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// storeRedis := redis_client.NewRedisStore(clientRedis)

	// initializing tokenManager to generate JWT

	tokenManager, err := jwt.NewManager(conf.JWTSecret, time.Duration(conf.AccessTokenTTL)*time.Minute, time.Duration(conf.RefreshTokenTTL)*time.Minute)
	if err != nil {
		logger.Fatal(err)
	}

	srv, err := server.NewHttpServer(ctx, logger, conf, confDB, redisConfig, conf.Port, tokenManager)

	if err != nil {

		return
	}

	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = srv.Shutdown(ctx); err != nil {
			fmt.Println("HTTP Server Shutdown")
		}
		close(stopped)
	}()

	logger.Info("Starting API Server...")

	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return
	}

	<-stopped

}
