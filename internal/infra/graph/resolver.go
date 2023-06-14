//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/dmcardoso/go-expert-challenge-clean-architecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}
