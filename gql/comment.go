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

func createCommentsField(commentType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Comments",
		Type: graphql.NewList(commentType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return httpClient.FetchComments()
		},
	}
}

func createCommentField(commentType *graphql.Object) *graphql.Field {
	return &graphql.Field{
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
	}
}
