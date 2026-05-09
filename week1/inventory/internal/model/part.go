package model

import "time"

type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
