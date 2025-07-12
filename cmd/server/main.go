package main

import (
	"log"
	"net/http"

	"github.com/Adit0507/news-feed-system/api"
	"github.com/Adit0507/news-feed-system/config"
	"github.com/Adit0507/news-feed-system/services"
	"github.com/Adit0507/news-feed-system/storage"
)

func main() {
	cfg := config.LoadConfig()
	store := storage.NewMemoryStore()
	userService := services.NewUserService(store)
	feedService := services.NewFeedService(store)
	router := api.NewRouter(userService, feedService)

	log.Printf("Server starting on %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
