package models

import (
	"time"
)

type Transaction struct {
	ID int
	UserID int
	amount float64
	Type string
	Category string
	description string
	date time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}