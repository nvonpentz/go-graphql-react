package resolvers

import (
	"github.com/nvonpentz/go-graphql-react/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service *services.Service
}

func NewWithDefaults() (*Resolver, error) {
	service, err := services.NewWithDefaults()
	if err != nil {
		return &Resolver{}, err
	}

	return &Resolver{
		Service: service,
	}, nil
}
