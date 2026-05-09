package v1

import (
	"context"
	desc "shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, orderUUID string, method int32) (string, error) {
	protoMethod := desc.PaymentMethod(method)

	resp, err := c.api.PayOrder(ctx, &desc.PayOrderRequest{
		OrderUuid:     orderUUID,
		PaymentMethod: protoMethod,
	})

	if err != nil {
		return "", err
	}

	return resp.GetTransactionUuid(), nil
}
