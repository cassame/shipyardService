package part

import (
	"inventory/internal/model"
	"inventory/internal/repository"
)

func NewRepository() repository.PartRepository {
	r := &repo{
		data: make(map[string]*model.Part),
	}
	r.setupMockData()
	return r
}

func (r *repo) setupMockData() {
	part := &model.Part{
		UUID:          "11111111-1111-1111-1111-111111111111",
		Name:          "Двигатель Гипердрайва",
		Description:   "Ускоряет до гипера",
		Price:         999.99,
		StockQuantity: 5,
		Category:      1,
	}
	r.data[part.UUID] = part
}
