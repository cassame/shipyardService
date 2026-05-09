package order

import (
	"context"
	"order/internal/model"
)

func (r *repo) Create(ctx context.Context, order *model.Order) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.data[order.UUID] = order

	return nil
}
