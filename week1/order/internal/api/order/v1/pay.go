package v1

import (
	"context"
	desc "shared/pkg/proto/order/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) PayOrder(ctx context.Context, req *desc.PayOrderRequest) (*desc.PayOrderResponse, error) {
	err := i.svc.PayOrder(ctx, req.GetOrderUuid(), int32(req.GetPaymentMethod()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "payment failed: %v", err)
	}

	return &desc.PayOrderResponse{Success: true}, nil
}
