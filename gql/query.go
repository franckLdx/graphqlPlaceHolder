package gql

import (
	"github.com/graphql-go/graphql"
)

func getQuery() *graphql.Object {
	userType := createUserType()
	postType := createPostType()
	commentType := createCommentType()
	albumType := createAlbumType()
	updateUserType(userType, postType, albumType)
	updatePostType(postType, commentType, userType)
	updateCommentType(commentType, postType)
	updateAlbumType(albumType, userType)

	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users":    createUsersField(userType),
			"user":     createUserField(userType),
			"posts":    createPostsField(postType),
			"post":     createPostField(postType),
			"comments": createCommentsField(commentType),
			"comment":  createCommentField(commentType),
			"albums":   createAlbumsField(albumType),
			"album":    createAlbumField(albumType),
		},
	}

	return graphql.NewObject(config)
}
