package order

import (
	"context"
	"encoding/json"
	"order/internal/model"
)

func (r *repo) Get(ctx context.Context, uuid string) (*model.Order, error) {

	query := `SELECT uuid, useruuid, status, total_price, items, created_at, updated_at
				FROM orders
				WHERE uuid = $1
	`

	order := &model.Order{}
	var itemsRaw []byte

	err := r.db.QueryRow(ctx, query, uuid).Scan(
		&order.UUID,
		&order.UserUUID,
		&order.Status,
		&order.TotalPrice,
		&itemsRaw,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(itemsRaw, &order.Items); err != nil {
		return nil, err
	}

	return order, nil
}
