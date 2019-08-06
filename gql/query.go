package gql

import (
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func getQuery() *graphql.Object {
	postType := createPostType()
	commentType := createCommentType()
	updatePostType(postType, commentType)
	updateCommentType(commentType, postType)

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
			"post": &graphql.Field{
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
			},
			"comments": &graphql.Field{
				Name: "Comments",
				Type: graphql.NewList(commentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return httpClient.FetchComments()
				},
			},
			"comment": &graphql.Field{
				Name: "Comment",
				Type: commentType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					return httpClient.FetchComment(id)
				},
			},
		},
	}
	return graphql.NewObject(config)
}
