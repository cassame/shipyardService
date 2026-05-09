package v1

import (
	"context"
	desc "shared/pkg/proto/inventory/v1"
)

func (c *client) GetPart(ctx context.Context, uuid string) (*desc.Part, error) {
	res, err := c.api.GetPart(ctx, &desc.GetPartRequest{
		Uuid: uuid,
	})
	if err != nil {
		return nil, err
	}

	return res.GetPart(), nil
}
