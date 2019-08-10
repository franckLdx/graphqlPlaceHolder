package gql

import (
	"github.com/graphql-go/graphql"
)

func getQuery() *graphql.Object {
	userType := createUserType()
	postType := createPostType()
	commentType := createCommentType()
	updateUserType(userType, postType)
	updatePostType(postType, commentType, userType)
	updateCommentType(commentType, postType)

	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users":    createUsersField(userType),
			"user":     createUserField(userType),
			"posts":    createPostsField(postType),
			"post":     createPostField(postType),
			"comments": createCommentsField(commentType),
			"comment":  createCommentField(commentType),
		},
	}

	return graphql.NewObject(config)
}
