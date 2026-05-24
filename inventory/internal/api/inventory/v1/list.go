package v1

import (
	"context"
	"inventory/internal/model"
	desc "shared/pkg/proto/inventory/v1"
)

func (s *Implementation) ListParts(ctx context.Context, req *desc.ListPartsRequest) (*desc.ListPartsResponse, error) {
	var filterModel *model.PartsFilter
	if req.GetFilter() != nil {
		f := req.GetFilter()

		categories := make([]int32, len(f.GetCategories()))
		for idx, cat := range f.GetCategories() {
			categories[idx] = int32(cat)
		}

		filterModel = &model.PartsFilter{
			UUIDs:                 f.GetUuids(),
			Names:                 f.GetNames(),
			Categories:            categories,
			ManufacturerCountries: f.GetManufacturerCountries(),
			Tags:                  f.GetTags(),
		}
	}

	partsModel, err := s.svc.ListParts(ctx, filterModel)
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
