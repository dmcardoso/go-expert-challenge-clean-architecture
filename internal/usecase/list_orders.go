package usecase

import (
	"github.com/dmcardoso/go-expert-challenge-clean-architecture/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	ordersResult, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var orders []OrderOutputDTO

	for _, order := range ordersResult {
		orders = append(orders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orders, nil
}
