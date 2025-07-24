package fraudclient

import (
	"context"
	"hexabank/api/proto/fraud"
	"hexabank/services/payment/domain/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FraudClient struct {
	client fraud.FraudServiceClient
}

func NewFraudClient(address string) (*FraudClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}

	client := fraud.NewFraudServiceClient(conn)
	return &FraudClient{
		client: client,
	}, nil
}

func (c *FraudClient) ValidatePayment(ctx context.Context, payment *model.Payment) (bool, error) {
	request := &fraud.PaymentRequest{
		Id:     payment.ID.String(),
		Amount: int32(payment.Amount),
	}

	response, err := c.client.FraudCheckHandler(ctx, request)
	if err != nil {
		return false, err
	}

	return response.IsFraudulent, nil
}
