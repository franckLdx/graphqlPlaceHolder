package gql

import (
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func getQuery() *graphql.Object {
	postType := createPostType()
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"posts": &graphql.Field{
				Name: "Posts",
				Type: graphql.NewList(postType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return httpClient.FetchPosts()
				},
			},
		},
	}
	return graphql.NewObject(config)
}

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
