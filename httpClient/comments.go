package httpClient

import (
	"fmt"
)

const commentsUrl = "comments"

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func FetchComments() (*[]Comment, error) {
	var comments []Comment
	if err := FetchList(commentsUrl, &comments); err != nil {
		return nil, fmt.Errorf("Failed to get comments list", err)
	}
	return &comments, nil
}

func FetchCommentsOfPosts(postId int) (*[]Comment, error) {
	filter := fmt.Sprint(postId, "/", commentsUrl)
	var comments []Comment
	if err := FetchList(filter, &comments); err != nil {
		return nil, fmt.Errorf("Failed to get comments of post", postId, err)
	}
	return &comments, nil
}
