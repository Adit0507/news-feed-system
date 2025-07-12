package api

import (
	"encoding/json"
	"net/http"

	"github.com/Adit0507/news-feed-system/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

func (h *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FollowerID string `json:"follower_id"`
		FolloweeID string `json:"followee_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	followerID, err := uuid.Parse(req.FollowerID)
	if err != nil {
		http.Error(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}

	followeeID, err := uuid.Parse(req.FolloweeID)
	if err != nil {
		http.Error(w, "Invalid followee ID", http.StatusBadRequest)
		return
	}

	if err := h.userService.Follow(followerID, followeeID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID  string `json:"user_id"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	post, err := h.feedService.CreatePost(userId, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) GetFeed(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	feed, err := h.feedService.GetFeed(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}
