package service

import (
	"context"
	"hexabank/services/fraud/domain/utils"
)

type FraudService struct{}

func (s *FraudService) FraudCheck(ctx context.Context, amount int) (bool, error) {
	isFibonacci := utils.IsFibonacci(amount)
	return isFibonacci, nil
}
