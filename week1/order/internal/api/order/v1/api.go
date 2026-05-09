package v1

import (
	"order/internal/service"
	desc "shared/pkg/proto/order/v1"
)

type Implementation struct {
	desc.UnimplementedOrderServiceServer
	svc service.OrderService
}

func NewImplementation(s service.OrderService) *Implementation {
	return &Implementation{svc: s}
}
