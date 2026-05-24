package order

import (
	"order/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.OrderRepository {
	return &repo{db: db}
}
