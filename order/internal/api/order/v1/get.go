package v1

import (
	"context"
	"errors"
	"order/internal/converter"
	"order/internal/model"
	openapi "shared/pkg/openapi/order/v1"
)

func (i *Implementation) GetOrder(ctx context.Context, params openapi.GetOrderParams) (openapi.GetOrderRes, error) {
	orderModel, err := i.svc.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &openapi.GetOrderNotFound{}, nil
		}

		return nil, err
	}

	return converter.ToAPIFromOrderModel(orderModel), nil
}
