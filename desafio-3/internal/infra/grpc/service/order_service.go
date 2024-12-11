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

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	// s.ListOrders()
	// s.
	// s.

	// orders, err := s.
	// if err != nil {
	// 	return nil, err
	// }

	// var ordersResponse []*pb.Order

	// for _, order := range orders {
	// 	orderResponse := &pb.Order{
	// 		Id:         order.ID,
	// 		Price:      order.Price,
	// 		Tax:        order.Tax,
	// 		FinalPrice: order.FinalPrice,
	// 	}

	// 	ordersResponse = append(ordersResponse, orderResponse)
	// }

	return &pb.OrderList{}, nil
	// return &pb.OrderList{Orders: ordersResponse}, nil
}
