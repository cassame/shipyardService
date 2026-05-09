package order

import (
	"context"
	"fmt"
	"order/internal/model"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateOrder(ctx context.Context, order *model.Order) (string, error) {
	var totalPrice float64

	for _, item := range order.Items {
		part, err := s.clients.Inventory().GetPart(ctx, item.PartUUID)
		if err != nil {
			return "", fmt.Errorf("товар %s не найден: %v", item.PartUUID, err)
		}

		totalPrice += part.Price * float64(item.Quantity)

	}

	order.UUID = uuid.New().String()
	order.Status = "NEW"
	order.TotalPrice = totalPrice
	order.CreatedAt = time.Now()

	err := s.orderRepo.Create(ctx, order)
	if err != nil {
		return "", fmt.Errorf("не удалось сохранить заказ: %v", err)
	}

	return order.UUID, nil
}
