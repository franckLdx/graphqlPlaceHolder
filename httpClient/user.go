package httpClient

import (
	"fmt"
)

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

const usersFilter = "users"

func FetchUsers() (*[]User, error) {
	var users []User
	if err := fetch(usersFilter, &users); err != nil {
		return nil, fmt.Errorf("Failed to get users list: %v", err)
	}
	return &users, nil
}

func FetchUser(id int) (*User, error) {
	var user User
	filter := fmt.Sprintf("%s/%d", usersFilter, id)
	if err := fetch(filter, &user); err != nil {
		return nil, fmt.Errorf("Failed to get user %d: %v ", id, err)
	}
	return &user, nil
}

func FetchPostsOfUser(id int) (*[]Post, error) {
	var posts []Post
	filter := fmt.Sprintf("posts?userId=%d", id)
	if err := fetch(filter, &posts); err != nil {
		return nil, fmt.Errorf("Failed to get user %d: %v ", id, err)
	}
	return &posts, nil
}
