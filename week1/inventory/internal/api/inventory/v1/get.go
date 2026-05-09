package v1

import (
	"context"
	"inventory/internal/converter"
	"log"
	desc "shared/pkg/proto/inventory/v1"
)

func (s *Implementation) GetPart(ctx context.Context, req *desc.GetPartRequest) (*desc.GetPartResponse, error) {
	log.Printf("Запрос детали с UUID: %s", req.GetUuid())

	partModel, err := s.svc.GetPart(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &desc.GetPartResponse{
		Part: converter.ToPartFromService(partModel),
	}, nil
}
