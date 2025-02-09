package main

import (
	"log"
	"task12/internal/infrastructure/http"
	"task12/internal/infrastructure/storage"
	"task12/internal/usecase/event"
	"task12/cmd/config"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}
	repo := storage.New()
	useCase := event.New(repo)
	server := http.NewServer(useCase, cfg.Port)
	server.StartServer()
}