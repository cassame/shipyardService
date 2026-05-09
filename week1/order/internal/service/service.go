package service

import (
	"context"
	"order/internal/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *model.Order) (string, error)
	GetOrder(ctx context.Context, orderUUID string) (*model.Order, error)
	CancelOrder(ctx context.Context, orderUUID string) error
	PayOrder(ctx context.Context, orderUUID string, paymentMethod int32) error
}
