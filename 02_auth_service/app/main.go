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
	"github.com/gMerl1n/notes_microservices/app/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
	"github.com/gMerl1n/notes_microservices/app/server"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	// initializing logger
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	// load .env variables

	if err := godotenv.Load("/home/username/Рабочий стол/notes_microservices/02_auth_service/.env"); err != nil {
		logger.Fatal(err)
	}

	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("Failed to load config")
	}

	tokenManager, err := jwt.NewManager(
		config.Token.JWTsecret,
		time.Duration(config.Token.AccessTokenTTL)*time.Minute,
		time.Duration(config.Token.RefreshTokenTTL)*time.Minute,
	)

	if err != nil {
		logger.Fatal(err)
	}

	ctx := context.Background()

	validate := validator.New(validator.WithRequiredStructEnabled())

	srv, err := server.NewHttpServer(ctx, logger, config, tokenManager, validate)

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
		logger.Fatal(err)
		return
	}

	<-stopped

}
