package httpClient

import (
	"fmt"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func FetchPosts() (*[]Post, error) {
	var posts []Post
	if err := FetchList("posts", &posts); err != nil {
		return nil, fmt.Errorf("Failed to get posts list", err)
	}
	return &posts, nil
}
