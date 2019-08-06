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

const commentsFilter = "comments"

func FetchComments() (*[]Comment, error) {
	var comments []Comment
	if err := fetch(commentsFilter, &comments); err != nil {
		return nil, fmt.Errorf("Failed to get comments list", err)
	}
	return &comments, nil
}

func FetchComment(id int) (*Comment, error) {
	filter := fmt.Sprintf("%s/%d", commentsFilter, id)
	var comments Comment
	if err := fetch(filter, &comments); err != nil {
		return nil, fmt.Errorf("Failed to get comment %d: %v", id, err)
	}
	return &comments, nil
}

func FetchCommentsOfPosts(postId int) (*[]Comment, error) {
	filter := fmt.Sprintf("posts/%d/%s", postId, commentsFilter)
	var comments []Comment
	if err := fetch(filter, &comments); err != nil {
		return nil, fmt.Errorf("Failed to get comments of post", postId, err)
	}
	return &comments, nil
}
