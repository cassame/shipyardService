package v1

import (
	"context"
	"errors"
	"order/internal/model"
	openapi "shared/pkg/openapi/order/v1"
)

func (i *Implementation) CancelOrder(ctx context.Context, params openapi.CancelOrderParams) (openapi.CancelOrderRes, error) {
	err := i.svc.CancelOrder(ctx, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &openapi.CancelOrderNotFound{}, nil
		}

		if errors.Is(err, model.ErrOrderAlreadyPaid) {
			return &openapi.CancelOrderConflict{}, nil
		}

		return nil, err
	}

	return &openapi.CancelOrderNoContent{}, nil
}
