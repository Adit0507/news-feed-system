package services

import (
	"errors"
	"time"

	"github.com/Adit0507/news-feed-system/models"
	"github.com/Adit0507/news-feed-system/storage"
	"github.com/Adit0507/news-feed-system/utils"
	"github.com/google/uuid"
)

type FeedService struct {
	store *storage.MemoryStore
}

func NewFeedService(store *storage.MemoryStore) *FeedService {
	return &FeedService{store: store}
}

func (s *FeedService) CreatePost(userId uuid.UUID, content string) (models.Post, error) {
	if _, exists := s.store.GetUser(userId); !exists {
		return models.Post{}, errors.New("user not found")
	}

	post := models.Post{
		ID:        uuid.New(),
		UserID:    userId,
		Content:   content,
		CreatedAT: time.Now(),
	}
	s.store.AddPost(post)

	// fan out on write for users with <1000 followers
	user, _ := s.store.GetUser(userId)
	if len(user.Followers) < 1000 {
		for followerID := range user.Followers {
			s.store.AddToFeed(followerID, post.ID)
		}
	}

	return post, nil
}

func (s *FeedService) GetFeed(userID uuid.UUID) ([]models.Post, error) {
	if _, exists := s.store.GetUser(userID); !exists {
		return nil, errors.New("user not found")
	}

	// fanout on write (precomputed feed)
	feed := s.store.GetFeed(userID)
	if len(feed) > 0 {
		return utils.RankPosts(feed), nil
	}

	// fallback to fanout on reead
	followees := s.store.GetFollowees(userID)
	var posts []models.Post
	for _, followeeID := range followees {
		userPosts := s.store.GetPostsByUser(followeeID)
		posts = append(posts, userPosts...)
	}

	return utils.RankPosts(posts), nil
}
