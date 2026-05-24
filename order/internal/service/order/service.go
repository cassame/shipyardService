package order

import (
	"order/internal/client/grpc"
	"order/internal/repository"
)

type Service struct {
	orderRepo repository.OrderRepository

	clients grpc.Clients
}

func NewService(
	orderRepo repository.OrderRepository,
	clients grpc.Clients,
) *Service {
	return &Service{
		orderRepo: orderRepo,
		clients:   clients,
	}
}
