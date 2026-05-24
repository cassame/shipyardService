package v1

import (
	"context"
	"order/internal/converter"
	openapi "shared/pkg/openapi/order/v1"
)

func (i *Implementation) CreateOrder(ctx context.Context, req *openapi.CreateOrderRequest) (*openapi.CreateOrderResponse, error) {
	orderModel := converter.ToOrderFromAPI(req)

	orderUUID, err := i.svc.CreateOrder(ctx, orderModel)
	if err != nil {
		return nil, err
	}

	return &openapi.CreateOrderResponse{
		UUID: orderUUID,
	}, nil
}
