package converter

import (
	"inventory/internal/model"
	desc "shared/pkg/proto/inventory/v1"
)

func ToPartFromService(part *model.Part) *desc.Part {
	if part == nil {
		return nil
	}

	return &desc.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      desc.Category(part.Category),
	}
}
