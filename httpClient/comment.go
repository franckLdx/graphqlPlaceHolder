package httpClient

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

const CommentResource Resource = "comments"

func FetchComments() (*[]Comment, error) {
	var comments []Comment
	err := FetchResources(CommentResource, &comments)
	return &comments, err
}

func FetchComment(commentId int) (*Comment, error) {
	var comment Comment
	err := FetchResource(CommentResource, commentId, &comment)
	return &comment, err
}

func FetchCommentsOfPosts(postId int) (*[]Comment, error) {
	var comments []Comment
	err := FetchSubResources(PostResource, postId, CommentResource, &comments)
	return &comments, err
}
