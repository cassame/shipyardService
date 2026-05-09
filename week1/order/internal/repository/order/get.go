package order

import (
	"context"
	"errors"
	"order/internal/model"
)

func (r *repo) Get(ctx context.Context, uuid string) (*model.Order, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	order, ok := r.data[uuid]
	if !ok {
		return nil, errors.New("order not found")
	}
	return order, nil
}
