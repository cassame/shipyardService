package v1

import (
	"context"
	desc "shared/pkg/proto/payment/v1"
)

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUUID string, method int32) (string, error)
}

type client struct {
	api desc.PaymentServiceClient
}

func NewClient(api desc.PaymentServiceClient) PaymentClient {
	return &client{api: api}
}
