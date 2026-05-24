package v1

import (
	"order/internal/service"
	openapi "shared/pkg/openapi/order/v1"
)

type Implementation struct {
	svc service.OrderService
}

func NewImplementation(s service.OrderService) *Implementation {
	return &Implementation{svc: s}
}

var _ openapi.Handler = (*Implementation)(nil)
