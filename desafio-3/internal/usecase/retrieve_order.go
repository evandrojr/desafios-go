package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type RetrieveOrderInputDTO struct {
	ID string `json:"id"`
}

type RetrieveOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type RetrieveOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderRetrieve   events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewRetrieveOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *RetrieveOrderUseCase) ExecuteRetrieve(input RetrieveOrderInputDTO) (RetrieveOrderOutputDTO, error) {

	entity := entity.Order{ID: input.ID}
	order, err := c.OrderRepository.Retrieve(entity)
	if err != nil {
		return RetrieveOrderOutputDTO{}, err
	}

	out := RetrieveOrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	return out, nil
}
