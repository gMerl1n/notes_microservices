package main

import (
	"fmt"
	"log"

	"github.com/gMerl1n/notes_microservices/internal/config"
)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("Failed to init config")
	}

	fmt.Println(config.Server)

}
