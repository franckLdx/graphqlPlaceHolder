package gql

import (
	"github.com/graphql-go/graphql"
)

func getQuery() *graphql.Object {
	userType := createUserType()
	postType := createPostType()
	commentType := createCommentType()
	albumType := createAlbumType()
	photoType := createPhotoType()
	todoType := createTodoType()
	updateUserType(userType, postType, albumType, todoType)
	updatePostType(postType, commentType, userType)
	updateCommentType(commentType, postType)
	updateAlbumType(albumType, userType, photoType)
	updatePhotoType(photoType, albumType)
	updateTodoType(todoType, userType)

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
			"photos":   createPhotosField(photoType),
			"photo":    createPhotoField(photoType),
			"todos":    createTodosField(todoType),
			"todo":     createTodoField(todoType),
		},
	}

	return graphql.NewObject(config)
}
