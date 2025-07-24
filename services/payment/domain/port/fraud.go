package port

import (
	"context"

	"hexabank/services/payment/domain/model"
)

type FraudClient interface {
	ValidatePayment(ctx context.Context, payment *model.Payment) (bool, error)
}
