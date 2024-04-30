package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/iriskin77/notes_microservices/internal/auth"
	"github.com/iriskin77/notes_microservices/internal/config"
	"github.com/iriskin77/notes_microservices/pkg/db"
	"github.com/iriskin77/notes_microservices/pkg/logging"
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

	// return &ConfigPostgres{
	// 	Host:     host,
	// 	Port:     port,
	// 	User:     user,
	// 	Password: password,
	// 	NameDB:   db_name,
	// 	SSLMode:  sslmode}

	ctx := context.Background()
	client, err := db.NewPostgresDB(ctx, confDB)

	if err != nil {
		fmt.Println(err.Error())
	}

	// initializing server
	repo := auth.NewRepository(client, logger)
	service := auth.NewService(repo, logger)
	h := auth.NewHandler(service, logger)

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
