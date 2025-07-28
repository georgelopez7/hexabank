package service

import (
	"context"
	"hexabank/services/fraud/domain/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type FraudService struct {
	tracer trace.Tracer
}

func NewFraudService() *FraudService {
	return &FraudService{
		tracer: otel.Tracer("fraud-service"),
	}
}

func (s *FraudService) FraudCheck(ctx context.Context, amount int) (bool, error) {
	_, span := s.tracer.Start(ctx, "fraud-check", trace.WithAttributes(
		attribute.Int("payment.amount", int(amount)),
	))
	defer span.End()

	isFibonacci := utils.IsFibonacci(amount)
	return isFibonacci, nil
}
