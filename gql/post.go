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

func updatePostType(postType, commentType *graphql.Object) {
	postType.AddFieldConfig(
		"comments",
		&graphql.Field{
			Type: graphql.NewList(commentType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				post, ok := p.Source.(httpClient.Post)
				if !ok {
					return nil, fmt.Errorf("failed to get post id")
				}
				return httpClient.FetchCommentsOfPosts(post.ID)
			},
		},
	)
}
