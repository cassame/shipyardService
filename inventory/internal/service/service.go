package service

import (
	"context"
	"inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, uuid string) (*model.Part, error)
	ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error)
	UpdateStock(ctx context.Context, uuid string, delta int64) error
}
