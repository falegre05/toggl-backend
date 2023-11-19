package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/graphql-go/graphql-go-handler"
	_ "github.com/mattn/go-sqlite3"
	"github.com/togglhire/backend-homework/schema"
	"github.com/togglhire/backend-homework/security"
)

func main() {
	schema, err := schema.Get()
	if err != nil {
		log.Fatal(err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	http.Handle("/graphql", security.AuthenticateMiddleware(gqlHandler))
	// http.Handle("/graphql", gqlHandler)
	port := getPort()
	fmt.Printf("GraphQL server is running on http://localhost:%s/graphql\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}

	return "3000"
}
