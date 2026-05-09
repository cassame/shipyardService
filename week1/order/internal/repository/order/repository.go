package order

import (
	"order/internal/model"
	"sync"
)

type repo struct {
	data map[string]*model.Order
	mx   sync.RWMutex
}

func NewRepository() *repo {
	return &repo{
		data: make(map[string]*model.Order),
	}
}
