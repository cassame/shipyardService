package part

import (
	"context"
	"fmt"
	"inventory/internal/model"
)

func (s *Service) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	part, err := s.repo.GetPart(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("get part from repo: %w", err)
	}

	return part, nil
}
