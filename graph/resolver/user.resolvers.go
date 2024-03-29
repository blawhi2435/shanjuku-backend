package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"fmt"

	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/graph"
	"github.com/blawhi2435/shanjuku-backend/graph/mapper"
	"github.com/blawhi2435/shanjuku-backend/graph/model"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.RegisterPayload, error) {
	var response *model.RegisterPayload = &model.RegisterPayload{}

	user, err := r.AuthUsecasse.Register(ctx, &domain.User{
		Account:  input.Account,
		Password: input.Password,
	})
	if err != nil {
		return response, cerror.GetGQLError(ctx, err)
	}

	modelUser := mapper.MappingUserDomainToModel(user)

	response = &model.RegisterPayload{
		User:  modelUser,
		Token: user.Token,
	}

	return response, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	var response *model.LoginPayload = &model.LoginPayload{}

	user, err := r.AuthUsecasse.Login(ctx, input.Account, input.Password)
	if err != nil {
		return response, cerror.GetGQLError(ctx, err)
	}

	modelUser := mapper.MappingUserDomainToModel(user)

	response = &model.LoginPayload{
		User:  modelUser,
		Token: user.Token,
	}

	return response, nil
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, input model.LogoutInput) (*model.LogoutPayload, error) {
	var response *model.LogoutPayload = &model.LogoutPayload{}

	err := r.AuthUsecasse.Logout(ctx, input.Account)
	if err != nil {
		return response, cerror.GetGQLError(ctx, err)
	}

	response = &model.LogoutPayload{
		Success: true,
	}

	return response, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
