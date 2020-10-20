package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	models "github.com/nvonpentz/go-graphql-react/internal/models/db"
	graphql1 "github.com/nvonpentz/go-graphql-react/internal/models/graphql"
)

func (r *mutationResolver) SignUp(ctx context.Context, input graphql1.UserInput) (*models.User, error) {
	return r.Service.SignUp(ctx, input)
}

func (r *mutationResolver) LogIn(ctx context.Context, input graphql1.UserInput) (*models.User, error) {
	return r.Service.LogIn(ctx, input)
}

func (r *mutationResolver) LogOut(ctx context.Context) (*bool, error) {
	return r.Service.LogOut(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.Service.Postgres)
}

func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	return r.Service.GetCurrentUser(ctx)
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.UpdatedAt.String(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
