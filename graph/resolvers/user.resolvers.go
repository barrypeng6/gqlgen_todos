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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   fmt.Sprintf("user_%d", rand.Int()),
		Name: input.Name,
	}
	r.MUsers = append(r.MUsers, user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context, first *int, after *string, last *int, before *string) (*model.UserConnection, error) {
	// check condition
	if err := helpers.CheckConnectionArgs(first, after, last, before); err != nil {
		return nil, err
	}

	var userEdges []*model.UserEdge
	if first != nil {
		for i, user := range r.MUsers {
			if i < *first {
				userEdges = append(userEdges, &model.UserEdge{
					Cursor: "",
					Node:   user,
				})
			}
		}
	}
	return &model.UserConnection{
		PageInfo: &model.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: userEdges,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	for _, _user := range r.MUsers {
		if _user.ID == id {
			user = _user
		}
	}
	return user, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, first *int, after *string, last *int, before *string) (*model.TodoConnection, error) {
	// check condition
	if err := helpers.CheckConnectionArgs(first, after, last, before); err != nil {
		return nil, err
	}

	// DB operation start
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
	} // DB operation end

	return &model.TodoConnection{
		PageInfo: &model.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: todoEdges,
	}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
