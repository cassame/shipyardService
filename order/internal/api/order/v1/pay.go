package v1

import (
	"context"
	"errors"
	"order/internal/model"
	openapi "shared/pkg/openapi/order/v1"

	"github.com/google/uuid"
)

func (i *Implementation) PayOrder(ctx context.Context, req *openapi.PayOrderRequest, params openapi.PayOrderParams) (openapi.PayOrderRes, error) {

	var methodInt int32
	switch req.PaymentMethod {
	case openapi.PaymentMethodCARD:
		methodInt = 1
	case openapi.PaymentMethodSBP:
		methodInt = 2
	case openapi.PaymentMethodCREDITCARD:
		methodInt = 3
	case openapi.PaymentMethodINVESTORMONEY:
		methodInt = 4
	default:
		methodInt = 0 // UNKNOWN
	}

	err := i.svc.PayOrder(ctx, params.OrderUUID, methodInt)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &openapi.PayOrderNotFound{}, nil
		}

		return nil, err
	}

	txUUID := uuid.New().String()

	return &openapi.PayOrderResponse{
		TransactionUUID: txUUID,
	}, nil
}
