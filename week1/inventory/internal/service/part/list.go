package part

import (
	"context"
	"inventory/internal/model"
)

func (s *Service) ListParts(ctx context.Context) ([]*model.Part, error) {
	//
	return s.repo.ListParts(ctx)
}
