package api

import (
	"github.com/Adit0507/news-feed-system/services"
	"github.com/go-chi/chi/v5"
)

func NewRouter(userService *services.UserService, feedService *services.FeedService) *chi.Mux {
	r := chi.NewRouter()
	h := NewHandler(userService, feedService)

	r.Post("/users", h.CreateUser)
	r.Post("/follow", h.Follow)
	r.Post("/posts", h.CreatePost)
	r.Get("/feed/{userID}", h.GetFeed)

	return r
}
