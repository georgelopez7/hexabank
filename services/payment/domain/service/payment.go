package service

import (
	"context"

	"hexabank/services/payment/domain/model"
	"hexabank/services/payment/domain/port"

	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepo port.PaymentRepository
}

func NewPaymentService(paymentRepo port.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, description string, amount int) (*model.Payment, error) {
	payment := model.NewPayment(description, amount)
	return payment, s.paymentRepo.CreatePayment(ctx, payment)
}

func (s *PaymentService) GetPayment(ctx context.Context, id uuid.UUID) (*model.Payment, error) {
	return s.paymentRepo.GetPayment(ctx, id)
}
