package gql

import (
	"fmt"
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func createPostType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	}
	return graphql.NewObject(config)
}

func updatePostType(postType, commentType, userType *graphql.Object) {
	postType.AddFieldConfig(
		"comments",
		&graphql.Field{
			Type: graphql.NewList(commentType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				post, err := getPost(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchCommentsOfPosts(post.ID)
			},
		},
	)
	postType.AddFieldConfig(
		"user",
		&graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				post, err := getPost(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchUser(post.UserID)
			},
		},
	)
}

func createPostsField(postType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Posts",
		Type: graphql.NewList(postType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return httpClient.FetchPosts()
		},
	}
}

func createPostField(postType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Post",
		Type: postType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return httpClient.FetchPost(id)
		},
	}
}

func getPost(p *graphql.ResolveParams) (*httpClient.Post, error) {
	if post, ok := p.Source.(httpClient.Post); ok {
		return &post, nil
	}
	if post, ok := p.Source.(*httpClient.Post); ok {
		return post, nil
	}
	return nil, fmt.Errorf("failed to get post")
}
