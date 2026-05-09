package part

import (
	"context"
	"inventory/internal/model"
)

func (r *repo) ListParts(ctx context.Context) ([]*model.Part, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	parts := make([]*model.Part, 0, len(r.data))
	for _, p := range r.data {
		parts = append(parts, p)
	}

	return parts, nil
}
