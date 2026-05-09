package model

import "time"

type Order struct {
	UUID       string
	UserUUID   string
	Items      []OrderItem // Список товаров
	TotalPrice float64
	Status     string // "NEW", "PAID", "CANCELLED"
	CreatedAt  time.Time
}

type OrderItem struct {
	PartUUID string
	Quantity int32
}
