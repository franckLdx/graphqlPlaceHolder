package httpClient

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const PostResource Resource = "posts"

func FetchPosts() (*[]Post, error) {
	var posts []Post
	err := FetchResources(PostResource, &posts)
	return &posts, err
}

func FetchPost(id int) (*Post, error) {
	var post Post
	err := FetchResource(PostResource, id, &post)
	return &post, err
}
