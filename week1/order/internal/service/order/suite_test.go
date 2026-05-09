package order

import (
	"order/internal/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrderServiceSuite struct {
	suite.Suite
	repoMock      *mocks.OrderRepository
	clientsMock   *mocks.Clients
	inventoryMock *mocks.InventoryClient
	paymentMock   *mocks.PaymentClient
	service       *Service
}

func (s *OrderServiceSuite) SetupTest() {
	s.repoMock = new(mocks.OrderRepository)
	s.clientsMock = new(mocks.Clients)
	s.inventoryMock = new(mocks.InventoryClient)
	s.paymentMock = new(mocks.PaymentClient)

	s.service = &Service{
		orderRepo: s.repoMock,
		clients:   s.clientsMock,
	}
}

func TestOrderServiceSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceSuite))
}
