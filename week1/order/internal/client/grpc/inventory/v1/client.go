package v1

import (
	"context"
	desc "shared/pkg/proto/inventory/v1"
)

type InventoryClient interface {
	GetPart(ctx context.Context, uuid string) (*desc.Part, error)
	//ListParts(ctx context.Context) (map[string]string, error)
}

type client struct {
	api desc.InventoryServiceClient
}

func NewClient(api desc.InventoryServiceClient) InventoryClient {
	return &client{api: api}
}
