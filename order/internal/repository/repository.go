package repository

import (
	"context"
	"order/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	Get(ctx context.Context, uuid string) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
}
