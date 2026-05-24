package v1

import (
	"inventory/internal/service"
	desc "shared/pkg/proto/inventory/v1"
)

type Implementation struct {
	desc.UnimplementedInventoryServiceServer
	svc service.InventoryService
}

func NewImplementation(s service.InventoryService) *Implementation {
	return &Implementation{
		svc: s,
	}
}
