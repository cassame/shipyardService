package v1

import (
	"context"
	"errors"
	"inventory/internal/model"
	desc "shared/pkg/proto/inventory/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateStock(ctx context.Context, req *desc.UpdateStockRequest) (*desc.UpdateStockResponse, error) {

	if req.GetUuid() == "" {
		return nil, status.Error(codes.InvalidArgument, "part uuid is required")
	}

	if req.GetDelta() == 0 {
		return nil, status.Error(codes.InvalidArgument, "delta cannot be zero")
	}

	err := i.svc.UpdateStock(ctx, req.GetUuid(), req.GetDelta())
	if err != nil {
		if errors.Is(err, model.ErrPartNotFound) {
			return nil, status.Error(codes.NotFound, "part not found")
		}

		if errors.Is(err, model.ErrInsufficientStock) {
			return nil, status.Error(codes.FailedPrecondition, "insufficient stock quantity")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &desc.UpdateStockResponse{}, nil
}
