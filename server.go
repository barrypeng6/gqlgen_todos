package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/barrypeng6/gqlgen-todos/graph/generated"
	"github.com/barrypeng6/gqlgen-todos/graph/model"
	"github.com/barrypeng6/gqlgen-todos/graph/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Mock data
	users := []*model.User{
		&model.User{
			ID:   "user_1234",
			Name: "Hello",
		},
	}
	todos := []*model.Todo{
		&model.Todo{
			ID:     "todo_1234",
			Text:   "Read books",
			Done:   false,
			UserID: "user_1234",
		},
	}
	resolvers := &resolvers.Resolver{
		MUsers: users,
		MTodos: todos,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
