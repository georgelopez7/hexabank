package service

import (
	"context"

	"hexabank/internal/errors"
	"hexabank/services/payment/domain/model"
	"hexabank/services/payment/domain/port"

	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepo port.PaymentRepository
	fraudClient port.FraudClient
}

func NewPaymentService(paymentRepo port.PaymentRepository, fraudClient port.FraudClient) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		fraudClient: fraudClient,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, description string, amount int) (*model.Payment, error) {
	payment := model.NewPayment(description, amount)

	isFraudulent, err := s.fraudClient.ValidatePayment(ctx, payment)
	if err != nil {
		return nil, errors.InternalError
	}

	if isFraudulent {
		return nil, errors.BadRequest
	}

	return payment, s.paymentRepo.CreatePayment(ctx, payment)
}

func (s *PaymentService) GetPayment(ctx context.Context, id uuid.UUID) (*model.Payment, error) {
	return s.paymentRepo.GetPayment(ctx, id)
}
