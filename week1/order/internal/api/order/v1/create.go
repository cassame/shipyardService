package v1

import (
	"context"
	"order/internal/converter"
	desc "shared/pkg/proto/order/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {

	orderModel := converter.ToOrderFromDesc(req)

	orderID, err := i.svc.CreateOrder(ctx, orderModel)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}

	return &desc.CreateOrderResponse{
		OrderUuid: orderID,
	}, nil
}
