package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/gMerl1n/notes_microservices/server"
	"github.com/go-playground/validator/v10"
)

func main() {

	// initializing logger
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("Failed to init config")
	}

	ctx := context.Background()

	validator := validator.New(validator.WithRequiredStructEnabled())

	tokenParser, err := jwt.NewTokenParser(config.Token.SigningKey)

	if err != nil {
		logger.Fatal("Failed to init Token Manager. ", err)
	}

	srv, err := server.NewHttpServer(ctx, logger, config, validator, tokenParser)

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
