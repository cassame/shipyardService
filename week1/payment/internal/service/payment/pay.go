package payment

import (
	"context"
	"payment/model"
	"time"

	"github.com/google/uuid"
)

func (s *Service) PayOrder(ctx context.Context, orderUUID string, method int32) (*model.Payment, error) {

	newPayment := model.Payment{
		TransactionUUID: uuid.New().String(),
		OrderUUID:       orderUUID,
		PaymentMethod:   method,
		Status:          "SUCCESS",
		CreatedAt:       time.Now(),
	}

	s.mx.Lock()
	s.data[newPayment.TransactionUUID] = newPayment
	s.mx.Unlock()

	return &newPayment, nil
}
