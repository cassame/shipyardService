package v1

import (
	"context"
	desc "shared/pkg/proto/payment/v1"
)

func (s *Implementation) PayOrder(ctx context.Context, req *desc.PayOrderRequest) (*desc.PayOrderResponse, error) {

	payment, err := s.svc.PayOrder(ctx, req.GetOrderUuid(), int32(req.GetPaymentMethod()))
	if err != nil {
		return nil, err
	}
	return &desc.PayOrderResponse{
		TransactionUuid: payment.TransactionUUID,
	}, nil
}
