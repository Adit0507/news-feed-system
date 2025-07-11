package models

import (
	"time"

	"github.com/google/uuid"
)

// system user
type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Followers map[uuid.UUID]bool
}

// user post
type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAT time.Time `json:"created_at"`
}

type Feed struct {
	UserID uuid.UUID `json:"user_id"`
	Posts  []Post    `json:"posts"`
}

type Relationship struct {	//follow relationship
	FollowerID uuid.UUID `json:"follower_id"`
	FolloweeID uuid.UUID `json:"followee_id"`
}
