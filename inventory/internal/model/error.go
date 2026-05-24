package model

import "errors"

var (
	// ErrPartNotFound возвращается, когда мы ищем деталь по UUID, а её нет в MongoDB
	ErrPartNotFound = errors.New("part not found")

	// ErrInsufficientStock возвращается, когда мы пытаемся списать деталей больше,
	// чем есть на складе (уход stock_quantity в минус)
	ErrInsufficientStock = errors.New("insufficient stock quantity")

	// ErrInvalidUUID возвращается, если передан некорректный формат UUID
	ErrInvalidUUID = errors.New("invalid uuid format")
)
