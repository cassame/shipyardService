package order

import (
	"context"
	"encoding/json"
	"order/internal/model"
)

func (r *repo) Create(ctx context.Context, order *model.Order) error {

	query := `
		INSERT INTO orders (uuid, useruuid, total_price, status, items, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	itemsJSON, _ := json.Marshal(order.Items)

	_, err := r.db.Exec(ctx, query,
		order.UUID,
		order.UserUUID,
		order.TotalPrice,
		order.Status,
		itemsJSON,
		order.CreatedAt,
		order.UpdatedAt,
	)

	return err
}
