package service

import (
	"context"
	"inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, uuid string) (*model.Part, error)
	ListParts(ctx context.Context) ([]*model.Part, error)
}
