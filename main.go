package main

import (
	"graphqlPlaceHolder/gql"
	"log"
	"net/http"
)

func main() {
	initLog()
	runServer()
}

func initLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

func runServer() {
	handler, err := gql.GetHandler()
	if err != nil {
		log.Fatal("failed to create graphql handler", err)
	}
	http.Handle("/graphql", handler)
	log.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
