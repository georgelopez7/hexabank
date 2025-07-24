package port

import "context"

type FraudService interface {
	FraudCheck(ctx context.Context, amount int) (bool, error)
}
