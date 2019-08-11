package gql

import (
	"fmt"
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func createUserType() *graphql.Object {
	geoCoordinatesType := createGeoCoordinatesType()
	addressType := createAddressType(geoCoordinatesType)
	companyType := createCompanyType()
	config := graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name":     &graphql.Field{Type: graphql.String},
			"address":  &graphql.Field{Type: addressType},
			"company":  &graphql.Field{Type: companyType},
			"email":    &graphql.Field{Type: graphql.String},
			"phone":    &graphql.Field{Type: graphql.String},
			"username": &graphql.Field{Type: graphql.String},
			"website":  &graphql.Field{Type: graphql.String},
		},
	}
	return graphql.NewObject(config)
}

func updateUserType(userType, postType, albumType *graphql.Object) {
	userType.AddFieldConfig(
		"posts",
		&graphql.Field{
			Type: graphql.NewList(postType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user, err := getUser(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchPostsOfUser(user.ID)
			},
		},
	)
	userType.AddFieldConfig(
		"albums",
		&graphql.Field{
			Type: graphql.NewList(postType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user, err := getUser(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchAlbumsOfUser(user.ID)
			},
		},
	)
}

func createUsersField(userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Users",
		Type: graphql.NewList(userType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return httpClient.FetchUsers()
		},
	}
}

func createUserField(userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "User",
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return httpClient.FetchUser(id)
		},
	}
}

func createGeoCoordinatesType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "GeoCoordinates",
		Fields: graphql.Fields{
			"lat": &graphql.Field{Type: graphql.String},
			"lng": &graphql.Field{Type: graphql.String},
		},
	}
	return graphql.NewObject(config)
}

func createAddressType(geoCoordinatesType *graphql.Object) *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"city":    &graphql.Field{Type: graphql.String},
			"geo":     &graphql.Field{Type: geoCoordinatesType},
			"street":  &graphql.Field{Type: graphql.String},
			"suite":   &graphql.Field{Type: graphql.String},
			"zipcode": &graphql.Field{Type: graphql.String},
		},
	}
	return graphql.NewObject(config)
}

func createCompanyType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"name": &graphql.Field{Type: graphql.String},
			"bs": &graphql.Field{
				Type: graphql.String,
			},
			"catchPhrase": &graphql.Field{Type: graphql.String},
		},
	}
	return graphql.NewObject(config)
}

func getUser(p *graphql.ResolveParams) (*httpClient.User, error) {
	if user, ok := p.Source.(httpClient.User); ok {
		return &user, nil
	}
	if user, ok := p.Source.(*httpClient.User); ok {
		return user, nil
	}
	return nil, fmt.Errorf("failed to get user")
}
