package model

import "time"

type Part struct {
	UUID          string    `bson:"_id"`
	Name          string    `bson:"name"`
	Description   string    `bson:"description"`
	Price         float64   `bson:"price"`
	StockQuantity int64     `bson:"stock_quantity"`
	Category      int32     `bson:"category"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}
