package main

import (
	"fmt"
	"log/slog"

	"github.com/iriskin77/notes_microservices/internal/config"
	"github.com/iriskin77/notes_microservices/internal/server"
	"github.com/iriskin77/notes_microservices/pkg/logging"
)

func main() {

	// initializing config
	conf := config.MustLoad()

	// initializing logger
	logger := logging.SetupLogger(conf.Env)
	// initializing server
	srv, err := server.NewHttpServer(conf.Port)

	if err != nil {
		fmt.Println(err.Error())
	}

	logger.Info("starting application", slog.String("env", conf.Env), slog.Any("cfg", conf))

	srv.ListenAndServe()

}
