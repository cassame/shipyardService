package v1

import (
	"context"
	desc "shared/pkg/proto/order/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*desc.CancelOrderResponse, error) {
	err := i.svc.CancelOrder(ctx, req.GetOrderUuid())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cancel order: %v", err)
	}

	return &desc.CancelOrderResponse{}, nil
}
