package v1

import (
	"context"
	"order/internal/converter"
	desc "shared/pkg/proto/order/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetOrder(ctx context.Context, req *desc.GetOrderRequest) (*desc.GetOrderResponse, error) {

	orderModel, err := i.svc.GetOrder(ctx, req.GetOrderUuid())
	if err != nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return converter.ToDescFromOrderModel(orderModel), nil
}
