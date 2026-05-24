package part

import "inventory/internal/repository"

type Service struct {
	repo repository.PartRepository
}

func NewService(r repository.PartRepository) *Service {
	return &Service{repo: r}
}
