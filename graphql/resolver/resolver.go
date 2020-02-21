package resolver

import (
	"context"

	graphql_server "github.com/alterra/graphql-server"
	"github.com/alterra/graphql-server/graphql/channel"
	model "github.com/alterra/graphql-server/graphql/models"
	"github.com/alterra/graphql-server/graphql/resolver/product"
	todo "github.com/alterra/graphql-server/graphql/resolver/todo"
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
	// gc, err := GinContextFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("context: " + fmt.Sprintf("%v", gc.Value("claims")))

	// Do something with claims
	// like checking customer id etc

	result, err := todo.CreateTodo(input)
	channel.PublishAll(result)
	return result, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todo.Todos()
}

func (r *queryResolver) GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	return product.GetAllProduct()
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
