package grpc

import (
	"context"
	"hexabank/api/proto/fraud"
	"hexabank/services/fraud/domain/port"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type FraudGRPC struct {
	fraud.UnimplementedFraudServiceServer
	fraudService port.FraudService
}

func NewFraudGRPC(fraudService port.FraudService) *FraudGRPC {
	return &FraudGRPC{
		fraudService: fraudService,
	}
}

func (s *FraudGRPC) FraudCheckHandler(ctx context.Context, req *fraud.PaymentRequest) (*fraud.FraudResponse, error) {
	var tracer = otel.Tracer("fraud-service/fraud-check-handler")

	paymentID := req.Id
	amount := req.Amount
	isFraudulent, err := s.fraudService.FraudCheck(ctx, int(amount))
	if err != nil {
		return nil, err
	}

	tracer.Start(ctx, "fraud-check-handler", trace.WithAttributes(
		attribute.String("payment.id", paymentID),
		attribute.Int("payment.amount", int(amount)),
	))

	if !isFraudulent {
		return &fraud.FraudResponse{
				IsFraudulent: isFraudulent,
				Message:      "Not Fraudulent",
			},
			nil
	}

	return &fraud.FraudResponse{
		IsFraudulent: isFraudulent,
		Message:      "Fraudulent",
	}, nil
}
