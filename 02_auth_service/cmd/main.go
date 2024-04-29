package main

import (
	"fmt"

	"github.com/iriskin77/notes_microservices/internal/config"
	"github.com/iriskin77/notes_microservices/internal/server"
)

func main() {

	// initializing config
	conf := config.MustLoad()

	// initializing logger

	// initializing server
	srv, err := server.NewHttpServer(conf.Port)

	if err != nil {
		fmt.Println(err.Error())
	}

	srv.ListenAndServe()

}
