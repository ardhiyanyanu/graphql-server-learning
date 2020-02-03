package resolver

import (
	"context"

	graphql_server "github.com/alterra/graphql-server"
	"github.com/alterra/graphql-server/channel"
	model "github.com/alterra/graphql-server/models"
	todo "github.com/alterra/graphql-server/resolver/todo"
	"github.com/google/uuid"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
}

func (r *Resolver) Mutation() graphql_server.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graphql_server.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() graphql_server.TodoResolver {
	return &todoResolver{r}
}
func (r *Resolver) Subscription() graphql_server.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	result, err := todo.CreateTodo(input)
	channel.PublishAll(result)
	return result, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todo.Todos()
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, todo *model.Todo) (*model.User, error) {
	return &model.User{
		ID:   todo.User.ID,
		Name: "Hei",
	}, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) TodoCreated(ctx context.Context) (<-chan *model.Todo, error) {
	uniqueKey := uuid.Must(uuid.NewRandom()).String()
	// Create new channel for request
	c := channel.NewChannel(uniqueKey)

	// Delete channel when done
	go func() {
		<-ctx.Done()
		channel.DeleteChannel(uniqueKey)
	}()

	return c, nil
}
