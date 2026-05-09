package converter

import (
	"order/internal/model"
	desc "shared/pkg/proto/inventory/v1"
)

func ToPartFromDesc(p *desc.Part) *model.OrderItem {
	if p == nil {
		return nil
	}
	return &model.OrderItem{
		PartUUID: p.Uuid,
		//map another fields for Order
	}
}
