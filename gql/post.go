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

func getPost(p *graphql.ResolveParams) (*httpClient.Post, error) {
	post, ok := p.Source.(*httpClient.Post)
	if !ok {
		return nil, fmt.Errorf("failed to get post")
	}
	return post, nil
}