package part

import (
	"context"
)

func (s *Service) UpdateStock(ctx context.Context, uuid string, delta int64) error {
	//business validation
	return s.repo.UpdateStock(ctx, uuid, delta)
}
