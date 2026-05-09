package order

import (
	"context"
	"fmt"
)

func (s *Service) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.orderRepo.Get(ctx, orderUUID)
	if err != nil {
		return fmt.Errorf("failed to get order: %w", err)
	}

	if order.Status == "PAID" {
		return fmt.Errorf("paid order cannot be cancelled")
	}

	order.Status = "CANCELLED"

	err = s.orderRepo.Update(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}

	return nil
}
