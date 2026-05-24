package service

import (
	"context"
	"payment/model"
)

type PaymentService interface {
	PayOrder(ctx context.Context, orderUUID string, method int32) (*model.Payment, error)
}
