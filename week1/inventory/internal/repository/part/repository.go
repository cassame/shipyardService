package part

import (
	"inventory/internal/model"
	"sync"
)

type repo struct {
	data map[string]*model.Part
	mx   sync.RWMutex
}
