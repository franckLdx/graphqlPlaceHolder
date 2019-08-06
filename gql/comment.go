package gql

import (
	"fmt"
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func createCommentType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	}
	return graphql.NewObject(config)
}

func updateCommentType(commentType, postType *graphql.Object) {
	commentType.AddFieldConfig(
		"post",
		&graphql.Field{
			Type: graphql.NewNonNull(postType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				comment, ok := p.Source.(httpClient.Comment)
				if !ok {
					return nil, fmt.Errorf("failed to get comment")
				}
				return httpClient.FetchPost(comment.PostID)
			},
		},
	)
}
