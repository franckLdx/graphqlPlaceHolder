package gql

import (
	"fmt"
	"graphqlPlaceHolder/httpClient"

	"github.com/graphql-go/graphql"
)

func createTodoType() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	}
	return graphql.NewObject(config)
}

func updateTodoType(todoType, userType *graphql.Object) {
	todoType.AddFieldConfig(
		"user",
		&graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todo, err := getTodo(&p)
				if err != nil {
					return nil, err
				}
				return httpClient.FetchUser(todo.UserID)
			},
		},
	)
}

func createTodosField(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Todos",
		Type: graphql.NewList(todoType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return httpClient.FetchTodos()
		},
	}
}

func createTodoField(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "Todo",
		Type: todoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return httpClient.FetchTodo(id)
		},
	}
}

func getTodo(p *graphql.ResolveParams) (*httpClient.Todo, error) {
	if todo, ok := p.Source.(httpClient.Todo); ok {
		return &todo, nil
	}
	if todo, ok := p.Source.(*httpClient.Todo); ok {
		return todo, nil
	}
	return nil, fmt.Errorf("failed to get todo")
}
