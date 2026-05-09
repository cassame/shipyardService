package order

import (
	"context"
	"errors"
	"order/internal/mocks"
	"order/internal/model"
	desc "shared/pkg/proto/inventory/v1"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateOrder(t *testing.T) {
	type fields struct {
		orderRepo       *mocks.OrderRepository
		clients         *mocks.Clients
		inventoryClient *mocks.InventoryClient
	}

	partUUID := "part-123"
	price := 100.0
	quantity := int32(2)

	tests := []struct {
		name    string
		order   *model.Order
		setup   func(f fields)
		wantErr bool
	}{
		{
			name: "Success: order created with correct price",
			order: &model.Order{
				Items: []model.OrderItem{{PartUUID: partUUID, Quantity: quantity}},
			},
			setup: func(f fields) {
				f.clients.On("Inventory").Return(f.inventoryClient)
				f.inventoryClient.On("GetPart", mock.Anything, partUUID).
					Return(&desc.Part{Price: price}, nil).Once()

				f.orderRepo.On("Create", mock.Anything, mock.MatchedBy(func(o *model.Order) bool {
					return o.TotalPrice == 200.0 && o.Status == "NEW"
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: item not found in inventory",
			order: &model.Order{
				Items: []model.OrderItem{{PartUUID: "unknown"}},
			},
			setup: func(f fields) {
				f.clients.On("Inventory").Return(f.inventoryClient)
				f.inventoryClient.On("GetPart", mock.Anything, "unknown").
					Return(nil, errors.New("not found")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				orderRepo:       new(mocks.OrderRepository),
				clients:         new(mocks.Clients),
				inventoryClient: new(mocks.InventoryClient),
			}
			tt.setup(f)

			s := &Service{
				orderRepo: f.orderRepo,
				clients:   f.clients,
			}

			uuid, err := s.CreateOrder(context.Background(), tt.order)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, uuid)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, uuid)
			}

			f.orderRepo.AssertExpectations(t)
			f.inventoryClient.AssertExpectations(t)
		})
	}
}
