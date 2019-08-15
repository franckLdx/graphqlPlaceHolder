package httpClient

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

const UserResource Resource = "users"

func FetchUsers() (*[]User, error) {
	var users []User
	err := FetchResources(UserResource, &users)
	return &users, err
}

func FetchUser(userId int) (*User, error) {
	var user User
	err := FetchResource(UserResource, userId, &user)
	return &user, err
}

func FetchPostsOfUser(userId int) (*[]Post, error) {
	var posts []Post
	err := FetchSubResources(UserResource, userId, PostResource, &posts)
	return &posts, err
}

func FetchAlbumsOfUser(userId int) (*[]Album, error) {
	var albums []Album
	err := FetchSubResources(UserResource, userId, AlbumResource, &albums)
	return &albums, err
}

func FetchTodosOfUser(userId int) (*[]Todo, error) {
	var todos []Todo
	err := FetchSubResources(UserResource, userId, TodoResource, &todos)
	return &todos, err
}
