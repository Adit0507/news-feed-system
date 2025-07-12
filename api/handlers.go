package api

import (
	"encoding/json"
	"net/http"

	"github.com/Adit0507/news-feed-system/services"
)

type Handler struct {
	userService *services.UserService
	feedService *services.FeedService
}

func NewHandler(userService *services.UserService, feedService *services.FeedService) *Handler {
	return &Handler{userService: userService, feedService: feedService}
}

// handles user creation
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err !=nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
