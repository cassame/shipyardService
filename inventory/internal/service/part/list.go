package part

import (
	"context"
	"inventory/internal/model"
)

func (s *Service) ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error) {
	//
	return s.repo.ListParts(ctx, filter)
}
