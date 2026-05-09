package v1

import (
	"payment/internal/service"
	desc "shared/pkg/proto/payment/v1"
)

type Implementation struct {
	desc.UnimplementedPaymentServiceServer
	svc service.PaymentService
}

func NewImplementation(s service.PaymentService) *Implementation {
	return &Implementation{
		svc: s,
	}
}
