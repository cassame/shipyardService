package grpc

import (
	inventory "order/internal/client/grpc/inventory/v1"
	payment "order/internal/client/grpc/payment/v1"
)

type Clients interface {
	Inventory() inventory.InventoryClient
	Payment() payment.PaymentClient
}

type clients struct {
	inventoryClient inventory.InventoryClient
	paymentClient   payment.PaymentClient
}

func NewClients(inv inventory.InventoryClient, pay payment.PaymentClient) Clients {
	return &clients{
		inventoryClient: inv,
		paymentClient:   pay,
	}
}

func (c *clients) Inventory() inventory.InventoryClient {
	return c.inventoryClient
}

func (c *clients) Payment() payment.PaymentClient {
	return c.paymentClient
}
