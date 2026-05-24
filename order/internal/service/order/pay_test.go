package order

import (
	"context"
	"errors"
	"order/internal/mocks"
	"order/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_PayOrder(t *testing.T) {
	type fields struct {
		orderRepo     *mocks.OrderRepository
		clients       *mocks.Clients
		paymentClient *mocks.PaymentClient
	}

	testUUID := "test-order-uuid"
	method := int32(1)

	tests := []struct {
		name    string
		setup   func(f fields)
		wantErr bool
	}{
		{
			name: "Success payment",
			setup: func(f fields) {
				order := &model.Order{UUID: testUUID, Status: "NEW"}
				f.orderRepo.On("Get", mock.Anything, testUUID).Return(order, nil).Once()

				f.clients.On("Payment").Return(f.paymentClient)
				f.paymentClient.On("PayOrder", mock.Anything, testUUID, method).
					Return("transaction-123", nil).Once()

				f.orderRepo.On("Update", mock.Anything, mock.MatchedBy(func(o *model.Order) bool {
					return o.Status == "PAID"
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: order already paid",
			setup: func(f fields) {
				order := &model.Order{UUID: testUUID, Status: "PAID"}
				f.orderRepo.On("Get", mock.Anything, testUUID).Return(order, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "Error: payment service failed",
			setup: func(f fields) {
				order := &model.Order{UUID: testUUID, Status: "NEW"}
				f.orderRepo.On("Get", mock.Anything, testUUID).Return(order, nil).Once()
				f.clients.On("Payment").Return(f.paymentClient)
				f.paymentClient.On("PayOrder", mock.Anything, testUUID, method).
					Return("", errors.New("bank rejected")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				orderRepo:     new(mocks.OrderRepository),
				clients:       new(mocks.Clients),
				paymentClient: new(mocks.PaymentClient),
			}
			tt.setup(f)

			s := &Service{
				orderRepo: f.orderRepo,
				clients:   f.clients,
			}

			err := s.PayOrder(context.Background(), testUUID, method)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			f.orderRepo.AssertExpectations(t)
			f.paymentClient.AssertExpectations(t)
		})
	}
}
