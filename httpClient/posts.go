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

const postsFilter = "posts"

func FetchPosts() (*[]Post, error) {
	var posts []Post
	if err := fetch(postsFilter, &posts); err != nil {
		return nil, fmt.Errorf("Failed to get posts list: %v", err)
	}
	return &posts, nil
}

func FetchPost(id int) (*Post, error) {
	var post Post
	filter := fmt.Sprintf("%s/%d", postsFilter, id)
	if err := fetch(filter, &post); err != nil {
		return nil, fmt.Errorf("Failed to get post %d: %v ", id, err)
	}
	return &post, nil
}
