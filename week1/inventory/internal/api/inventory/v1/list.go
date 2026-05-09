package v1

import (
	"context"
	desc "shared/pkg/proto/inventory/v1"
)

func (s *Implementation) ListParts(ctx context.Context, req *desc.ListPartsRequest) (*desc.ListPartsResponse, error) {
	partsModel, err := s.svc.ListParts(ctx)
	if err != nil {
		return nil, err
	}

	var protoParts []*desc.Part
	for _, p := range partsModel {
		protoParts = append(protoParts, &desc.Part{
			Uuid:          p.UUID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			StockQuantity: p.StockQuantity,
			Category:      desc.Category(p.Category),
		})
	}

	return &desc.ListPartsResponse{
		Parts: protoParts,
	}, nil
}
