package httpClient

import (
	"fmt"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func FetchComments() (*[]Comment, error) {
	var comments []Comment
	if err := FetchList("comments", &comments); err != nil {
		return nil, fmt.Errorf("Failed to get comments list", err)
	}
	return &comments, nil
}
