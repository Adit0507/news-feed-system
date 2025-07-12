package services

import (
	"errors"

	"github.com/Adit0507/news-feed-system/models"
	"github.com/Adit0507/news-feed-system/storage"
	"github.com/google/uuid"
)

type UserService struct {
	store *storage.MemoryStore
}

func NewUserService(store *storage.MemoryStore) *UserService {
	return &UserService{store: store}
}

// creates a new user
func (s *UserService) CreateUser(username string) (models.User, error) {
	user := models.User{
		ID:        uuid.New(),
		Username:  username,
		Followers: make(map[uuid.UUID]bool),
	}

	s.store.AddUser(user)
	return user, nil
}

// establishes a follow relationship
func (s *UserService) Follow(followerID, followeeID uuid.UUID) error {
	if _, exists := s.store.GetUser(followerID); !exists {
		return errors.New("follower not found")
	}
	if _, exists := s.store.GetUser(followeeID); !exists {
		return errors.New("followee not found")
	}

	s.store.AddRelationship(followerID, followeeID)

	return nil
}
