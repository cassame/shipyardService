package order

import (
	"context"
	"fmt"
	"order/internal/model"
)

func (s *Service) GetOrder(ctx context.Context, uuid string) (*model.Order, error) {
	order, err := s.orderRepo.Get(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get order from repository: %w", err)
	}
	return order, nil
}
