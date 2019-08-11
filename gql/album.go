package gql

import (
	"fmt"
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func createAlbumType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Album",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
		},
	}
	return graphql.NewObject(config)
}

func updateAlbumType(albumType, userType *graphql.Object) {
	albumType.AddFieldConfig(
		"user",
		&graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				album, err := getAlbum(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchUser(album.UserID)
			},
		},
	)
}

func createAlbumsField(commentType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Albums",
		Type: graphql.NewList(commentType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return httpClient.FetchAlbums()
		},
	}
}

func createAlbumField(commentType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Album",
		Type: commentType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return httpClient.FetchAlbum(id)
		},
	}
}

func getAlbum(p *graphql.ResolveParams) (*httpClient.Album, error) {
	if album, ok := p.Source.(httpClient.Album); ok {
		return &album, nil
	}
	if album, ok := p.Source.(*httpClient.Album); ok {
		return album, nil
	}
	return nil, fmt.Errorf("failed to get album")
}
