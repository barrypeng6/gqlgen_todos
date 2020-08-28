package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/barrypeng6/gqlgen-todos/graph/generated"
	"github.com/barrypeng6/gqlgen-todos/graph/helpers"
	"github.com/barrypeng6/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	hasUser := false
	for _, user := range r.MUsers {
		if user.ID == input.UserID {
			hasUser = true
		}
	}
	if hasUser {
		todo := &model.Todo{
			Text:   input.Text,
			ID:     fmt.Sprintf("todo_%d", rand.Int()),
			UserID: input.UserID,
		}
		r.MTodos = append(r.MTodos, todo)
		return todo, nil
	}
	return nil, fmt.Errorf("No this user (%s)", input.UserID)
}

func (r *queryResolver) Todos(ctx context.Context, first *int, after *string, last *int, before *string) (*model.TodoConnection, error) {
	// check condition
	if err := helpers.CheckConnectionArgs(first, after, last, before); err != nil {
		return nil, err
	}

	var todoEdges []*model.TodoEdge
	if first != nil {
		for i, todo := range r.MTodos {
			if i < *first {
				todoEdges = append(todoEdges, &model.TodoEdge{
					Cursor: "",
					Node:   todo,
				})
			}
		}
	}
	return &model.TodoConnection{
		PageInfo: &model.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: todoEdges,
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	var user *model.User
	for _, _user := range r.MUsers {
		if _user.ID == obj.UserID {
			user = _user
		}
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
