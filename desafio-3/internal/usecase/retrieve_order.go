package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type RetrievalOrderInputDTO struct {
	ID string `json:"id"`
}

type RetrievalOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type RetrievalOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderRetrieval   events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewRetrievalOrderUseCase(
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

func (c *RetrievalOrderUseCase) ExecuteRetrieval(input RetrievalOrderInputDTO) (RetrievalOrderOutputDTO, error) {

	entity := entity.Order{ID: input.ID}
	order, err := c.OrderRepository.Retrieve(entity)
	if err != nil {
		return RetrievalOrderOutputDTO{}, err
	}

	out := RetrievalOrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	return out, nil
}
