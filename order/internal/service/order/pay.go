package order

import (
	"context"
	"fmt"
)

func (s *Service) PayOrder(ctx context.Context, uuid string, method int32) error {
	order, err := s.orderRepo.Get(ctx, uuid)
	if err != nil {
		return fmt.Errorf("order not found: %w", err)
	}

	if order.Status == "PAID" {
		return fmt.Errorf("order is already paid")
	}
	if order.Status == "CANCELLED" {
		return fmt.Errorf("cannot pay for cancelled order")
	}

	transactionUUID, err := s.clients.Payment().PayOrder(ctx, uuid, method)
	if err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}
	fmt.Printf("Transaction successful: %s\n", transactionUUID)

	order.Status = "PAID"

	err = s.orderRepo.Update(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
}
