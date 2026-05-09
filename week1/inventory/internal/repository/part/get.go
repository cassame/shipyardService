package part

import (
	"context"
	"inventory/internal/model"
	"log"
)

func (r *repo) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	log.Printf("Запрос детали с UUID: %s", uuid)

	return &model.Part{
		UUID:          uuid,
		Name:          "Двигатель Гипердрайва",
		Description:   "Ускоряет до гипера",
		Price:         999.99,
		StockQuantity: 5,
		Category:      1,
	}, nil
}
