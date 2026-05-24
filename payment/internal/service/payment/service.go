package payment

import (
	"payment/model"
	"sync"
)

type Service struct {
	data map[string]model.Payment
	mx   sync.RWMutex
}

func NewService() *Service {
	return &Service{
		data: make(map[string]model.Payment),
	}
}
