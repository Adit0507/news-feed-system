package utils

import (
	"sort"

	"github.com/Adit0507/news-feed-system/models"
)

func RankPosts(posts []models.Post) []models.Post {
	sortedPosts := make([]models.Post, len(posts))
	copy(sortedPosts, posts)
	sort.Slice(sortedPosts, func(i, j int) bool {
		return sortedPosts[i].CreatedAT.After(sortedPosts[j].CreatedAT)
	})

	return sortedPosts
}
