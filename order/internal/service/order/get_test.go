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

func TestService_GetOrder(t *testing.T) {
	type fields struct {
		orderRepo *mocks.OrderRepository
	}
	type args struct {
		ctx  context.Context
		uuid string
	}

	testUUID := "test-order-uuid"
	expectedOrder := &model.Order{
		UUID:   testUUID,
		Status: "NEW",
	}

	tests := []struct {
		name    string
		setup   func(f fields)
		args    args
		want    *model.Order
		wantErr bool
	}{
		{
			name: "Success: order found",
			setup: func(f fields) {
				f.orderRepo.On("Get", mock.Anything, testUUID).
					Return(expectedOrder, nil).Once()
			},
			args: args{
				ctx:  context.Background(),
				uuid: testUUID,
			},
			want:    expectedOrder,
			wantErr: false,
		},
		{
			name: "Error: repository failure",
			setup: func(f fields) {
				f.orderRepo.On("Get", mock.Anything, "wrong-uuid").
					Return(nil, errors.New("db error")).Once()
			},
			args: args{
				ctx:  context.Background(),
				uuid: "wrong-uuid",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				orderRepo: &mocks.OrderRepository{},
			}
			tt.setup(f)

			s := &Service{
				orderRepo: f.orderRepo,
			}

			res, err := s.GetOrder(tt.args.ctx, tt.args.uuid)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, res)
			}

			f.orderRepo.AssertExpectations(t)
		})
	}

}
