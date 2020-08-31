package resolvers

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/barrypeng6/gqlgen-todos/graph/auth"
	"github.com/barrypeng6/gqlgen-todos/graph/generated"
	"github.com/barrypeng6/gqlgen-todos/graph/model"
	"github.com/labstack/echo/v4"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MTodos []*model.Todo
	MUsers []*model.User
}

func CreateRootResolver() generated.Config {

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

	c := generated.Config{
		Resolvers: &Resolver{
			MUsers: users,
			MTodos: todos,
		},
	}

	// Authentication
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ctxUserID := ctx.Value(auth.UserCtxKey)
		if ctxUserID == nil {
			return nil, echo.ErrUnauthorized
		}
		return next(ctx)
	}

	return c

}
