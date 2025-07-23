package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	Amount      int       `json:"amount" db:"amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func NewPayment(description string, amount int) *Payment {
	return &Payment{
		ID:          uuid.New(),
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
}
