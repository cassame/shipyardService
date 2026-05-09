package model

import "time"

type Payment struct {
	TransactionUUID string
	OrderUUID       string
	UserUUID        string
	PaymentMethod   int32
	Status          string
	CreatedAt       time.Time
}
