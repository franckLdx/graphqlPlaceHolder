package gql

import (
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func getQuery() *graphql.Object {
	postType := createPostType()
	commentType := createCommentType()
	updatePostType(postType, commentType)

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
			"comments": &graphql.Field{
				Name: "Comments",

				Type: graphql.NewList(commentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return httpClient.FetchComments()
				},
			},
		},
	}
	return graphql.NewObject(config)
}
