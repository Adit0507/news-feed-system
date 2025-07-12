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

// retrieves user by ID
func (s *MemoryStore) GetUser(userID uuid.UUID) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[userID]
	return user, exists
}

func (s *MemoryStore) AddPost(post models.Post) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.posts[post.ID] = post
}

func (s *MemoryStore) GetPostsByUser(userID uuid.UUID) []models.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var posts []models.Post
	for _, post := range s.posts {
		if post.UserID == userID {
			posts = append(posts, post)
		}
	}

	return posts
}

// fan out write- adds a post to user's feed
func (s *MemoryStore) AddToFeed(userID, postId uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[postId]
	if exists {
		s.feeds[userID] = append(s.feeds[userID], post)
	}
}

// retrivein a user computed feed
func (s *MemoryStore) GetFeed(userID uuid.UUID) []models.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	feed, exists := s.feeds[userID]
	if !exists {
		return []models.Post{}
	}

	return feed
}

// follow relationship
func (s *MemoryStore) AddRelationship(followerID, followeeID uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.relationships[followerID] = append(s.relationships[followerID], followeeID)
	user, exists := s.users[followerID]
	if exists {
		user.Followers[followeeID] = true
		s.users[followerID] = user
	}
}

func (s *MemoryStore) GetFollowees(followerID uuid.UUID) []uuid.UUID {
	s.mu.RLock()
	defer s.mu.RUnlock()

	followees, exists := s.relationships[followerID]
	if !exists {
		return []uuid.UUID{}
	}

	return followees
}
