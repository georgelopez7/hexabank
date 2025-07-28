package grpc

import (
	"context"
	"hexabank/api/proto/fraud"
	"hexabank/services/fraud/domain/port"
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

func (h *FraudGRPC) FraudCheckHandler(ctx context.Context, req *fraud.PaymentRequest) (*fraud.FraudResponse, error) {
	amount := req.Amount
	isFraudulent, err := h.fraudService.FraudCheck(ctx, int(amount))
	if err != nil {
		return nil, err
	}

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
