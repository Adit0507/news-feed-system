package storage

import (
	"sync"

	"github.com/Adit0507/news-feed-system/models"
	"github.com/google/uuid"
)

// simultaes a database and cache
type MemoryStore struct {
	users         map[uuid.UUID]models.User
	posts         map[uuid.UUID]models.Post
	feeds         map[uuid.UUID][]models.Post
	relationships map[uuid.UUID][]uuid.UUID
	mu            sync.RWMutex
}

// initializin the store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:         make(map[uuid.UUID]models.User),
		posts:         make(map[uuid.UUID]models.Post),
		feeds:         make(map[uuid.UUID][]models.Post),
		relationships: make(map[uuid.UUID][]uuid.UUID),
	}
}

func (s *MemoryStore) AddUser(user models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[user.ID] = user
}


