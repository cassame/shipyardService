package part

import (
	"context"
	"fmt"
	"inventory/internal/model"
)

func (s *Service) UpdateStock(ctx context.Context, uuid string, delta int64) error {

	part, err := s.repo.GetPart(ctx, uuid)
	if err != nil {
		return fmt.Errorf("get part from repo: %w", err)
	}

	if delta < 0 && (part.StockQuantity+delta) < 0 {
		return model.ErrInsufficientStock
	}

	err = s.repo.UpdateStock(ctx, uuid, delta)
	if err != nil {
		return fmt.Errorf("update stock in repo: %w", err)
	}

	return nil
}
