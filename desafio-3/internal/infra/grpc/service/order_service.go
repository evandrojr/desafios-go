package service

import (
	"context"

	"desafio3/internal/infra/grpc/pb"
	"desafio3/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.ExecuteCreate(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

// func (c *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
// 	categories, err := c.CategoryDB.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var categoriesResponse []*pb.Category

// 	for _, category := range categories {
// 		categoryResponse := &pb.Category{
// 			Id:          category.ID,
// 			Name:        category.Name,
// 			Description: category.Description,
// 		}

// 		categoriesResponse = append(categoriesResponse, categoryResponse)
// 	}

// 	return &pb.CategoryList{Categories: categoriesResponse}, nil
// }

// func (s *OrderService) ListOrder(ctx context.Context, in *pb.) (*pb.CreateOrderResponse, error) {
// 	&pb.

// 	// dto := usecase.OrderInputDTO{
// 	// 	ID:    in.Id,
// 	// 	Price: float64(in.Price),
// 	// 	Tax:   float64(in.Tax),
// 	// }
// 	// output, err := s.CreateOrderUseCase.ExecuteCreate(dto)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return &pb.CreateOrderResponse{
// 	// 	Id:         output.ID,
// 	// 	Price:      float32(output.Price),
// 	// 	Tax:        float32(output.Tax),
// 	// 	FinalPrice: float32(output.FinalPrice),
// 	// }, nil
// }
