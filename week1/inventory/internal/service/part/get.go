package part

import (
	"context"
	"inventory/internal/model"
)

func (s *Service) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	//
	return s.repo.GetPart(ctx, uuid)
}
