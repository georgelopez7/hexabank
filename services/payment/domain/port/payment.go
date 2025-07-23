package port

import (
	"context"

	"hexabank/payment/domain/model"

	"github.com/google/uuid"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, description string, amount int) (*model.Payment, error)
	GetPayment(ctx context.Context, id uuid.UUID) (*model.Payment, error)
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	GetPayment(ctx context.Context, id uuid.UUID) (*model.Payment, error)
}
