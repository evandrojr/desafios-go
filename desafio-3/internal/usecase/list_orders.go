package usecase

import (
	"desafio3/internal/entity"
	"desafio3/pkg/events"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderList       events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	// OrderCreated events.EventInterface,
	// EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		// OrderCreated:    OrderCreated,
		// EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) ExecuteList() ([]ListOrderOutputDTO, error) {

	var listOrderOutputDTOList []ListOrderOutputDTO
	orders, err := c.OrderRepository.List()
	if err != nil {
		return listOrderOutputDTOList, err
	}

	for _, order := range orders {

		orderDto := ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		listOrderOutputDTOList = append(listOrderOutputDTOList, orderDto)

	}

	return listOrderOutputDTOList, nil
}
