package fraudclient

import (
	"context"
	"hexabank/services/payment/domain/model"

	"github.com/stretchr/testify/mock"
)

type MockFraudClient struct {
	mock.Mock
}

func (m *MockFraudClient) ValidatePayment(ctx context.Context, payment *model.Payment) (bool, error) {
	args := m.Called(ctx, payment)
	return args.Bool(0), args.Error(1)
}
