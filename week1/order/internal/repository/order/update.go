package order

import (
	"context"
	"errors"
	"order/internal/model"
)

func (r *repo) Update(ctx context.Context, order *model.Order) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.data[order.UUID]; !ok {
		return errors.New("order not found")
	}

	r.data[order.UUID] = order
	return nil
}
