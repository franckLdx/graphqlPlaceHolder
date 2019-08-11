package gql

import (
	"fmt"
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func createPhotoType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Photo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
			"ThumbnailURL": &graphql.Field{
				Type: graphql.String,
			},
		},
	}
	return graphql.NewObject(config)
}

func updatePhotoType(photoType, albumType *graphql.Object) {
	photoType.AddFieldConfig(
		"album",
		&graphql.Field{
			Type: albumType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				photo, err := getPhoto(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchAlbum(photo.AlbumId)
			},
		},
	)
}

func createPhotosField(photoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Photos",
		Type: graphql.NewList(photoType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return httpClient.FetchPhotos()
		},
	}
}

func createPhotoField(photoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Photo",
		Type: photoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return httpClient.FetchPhoto(id)
		},
	}
}

func getPhoto(p *graphql.ResolveParams) (*httpClient.Photo, error) {
	if photo, ok := p.Source.(httpClient.Photo); ok {
		return &photo, nil
	}
	if photo, ok := p.Source.(*httpClient.Photo); ok {
		return photo, nil
	}
	return nil, fmt.Errorf("failed to get photo")
}
