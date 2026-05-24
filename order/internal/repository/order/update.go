package order

import (
	"context"
	"encoding/json"
	"errors"
	"order/internal/model"
)

func (r *repo) Update(ctx context.Context, order *model.Order) error {

	query := `
		UPDATE orders
		SET status = $2, total_price = $3, items = $4, updated_at = NOW()
		WHERE uuid = $1
	`

	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	res, err := r.db.Exec(ctx, query,
		order.UUID,
		order.Status,
		order.TotalPrice,
		itemsJSON,
	)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("order not found, nothing to update")
	}

	return nil
}
