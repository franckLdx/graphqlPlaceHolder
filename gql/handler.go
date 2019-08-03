package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func GetHandler() (*handler.Handler, error) {
	schema, err := getSchema()
	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	}), nil
}

func getSchema() (*graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: getQuery(),
	})
	if err != nil {
		return nil, err
	}
	return &schema, nil
}
