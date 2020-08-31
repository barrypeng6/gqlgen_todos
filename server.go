package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/barrypeng6/gqlgen-todos/graph/auth"
	"github.com/barrypeng6/gqlgen-todos/graph/generated"
	"github.com/barrypeng6/gqlgen-todos/graph/resolvers"
	"github.com/labstack/echo/v4"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	// e.Use(middleware.Gzip())
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("my_secret"),
	// }))

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Any("/", echo.WrapHandler(playground.Handler("GraphQL Playground", "/query")))
	e.Any("/query", echo.WrapHandler(auth.Middleware(handler.NewDefaultServer(generated.NewExecutableSchema(resolvers.CreateRootResolver())))))

	e.Logger.Fatal(e.Start(":" + port))
}
