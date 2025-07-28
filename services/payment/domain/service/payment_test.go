package service

import (
	"context"
	"errors"
	"testing"

	fraudclient "hexabank/services/payment/adapters/fraud-client"
	"hexabank/services/payment/adapters/postgres"
	"hexabank/services/payment/domain/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PaymentService_CreatePayment(t *testing.T) {
	t.Run("should create payment successfully", func(t *testing.T) {
		mockPaymentRepo := new(postgres.MockPaymentRepo)
		mockFraudClient := new(fraudclient.MockFraudClient)
		paymentService := NewPaymentService(mockPaymentRepo, mockFraudClient)

		description := "test payment"
		amount := 100

		mockFraudClient.On("ValidatePayment", mock.Anything, mock.AnythingOfType("*model.Payment")).Return(false, nil)
		mockPaymentRepo.On("CreatePayment", mock.Anything, mock.AnythingOfType("*model.Payment")).Return(nil)

		createdPayment, err := paymentService.CreatePayment(context.Background(), description, amount)

		assert.NoError(t, err)
		assert.NotNil(t, createdPayment)
		assert.Equal(t, description, createdPayment.Description)
		assert.Equal(t, amount, createdPayment.Amount)
		mockPaymentRepo.AssertExpectations(t)
		mockFraudClient.AssertExpectations(t)
	})

	t.Run("should return error when fraud is detected", func(t *testing.T) {
		mockPaymentRepo := new(postgres.MockPaymentRepo)
		mockFraudClient := new(fraudclient.MockFraudClient)
		paymentService := NewPaymentService(mockPaymentRepo, mockFraudClient)

		description := "fraudulent payment"
		amount := 200

		mockFraudClient.On("ValidatePayment", mock.Anything, mock.AnythingOfType("*model.Payment")).Return(true, nil)

		createdPayment, err := paymentService.CreatePayment(context.Background(), description, amount)

		assert.Error(t, err)
		assert.Nil(t, createdPayment)
		mockFraudClient.AssertExpectations(t)
		mockPaymentRepo.AssertNotCalled(t, "CreatePayment")
	})

	t.Run("should return error when fraud check fails", func(t *testing.T) {
		mockPaymentRepo := new(postgres.MockPaymentRepo)
		mockFraudClient := new(fraudclient.MockFraudClient)
		paymentService := NewPaymentService(mockPaymentRepo, mockFraudClient)

		description := "payment with error"
		amount := 300

		mockFraudClient.On("ValidatePayment", mock.Anything, mock.AnythingOfType("*model.Payment")).Return(false, errors.New("fraud client error"))

		createdPayment, err := paymentService.CreatePayment(context.Background(), description, amount)

		assert.Error(t, err)
		assert.Nil(t, createdPayment)
		mockFraudClient.AssertExpectations(t)
		mockPaymentRepo.AssertNotCalled(t, "CreatePayment")
	})
}

func Test_PaymentService_GetPayment(t *testing.T) {
	t.Run("should get payment successfully", func(t *testing.T) {
		mockPaymentRepo := new(postgres.MockPaymentRepo)
		mockFraudClient := new(fraudclient.MockFraudClient)
		paymentService := NewPaymentService(mockPaymentRepo, mockFraudClient)

		paymentID := uuid.New()
		payment := &model.Payment{ID: paymentID, Description: "test payment", Amount: 100}

		mockPaymentRepo.On("GetPayment", mock.Anything, paymentID).Return(payment, nil)

		retrievedPayment, err := paymentService.GetPayment(context.Background(), paymentID)

		assert.NoError(t, err)
		assert.NotNil(t, retrievedPayment)
		assert.Equal(t, paymentID, retrievedPayment.ID)
		mockPaymentRepo.AssertExpectations(t)
	})

	t.Run("should return error when payment not found", func(t *testing.T) {
		mockPaymentRepo := new(postgres.MockPaymentRepo)
		mockFraudClient := new(fraudclient.MockFraudClient)
		paymentService := NewPaymentService(mockPaymentRepo, mockFraudClient)

		paymentID := uuid.New()

		mockPaymentRepo.On("GetPayment", mock.Anything, paymentID).Return(nil, errors.New("payment not found"))

		retrievedPayment, err := paymentService.GetPayment(context.Background(), paymentID)

		assert.Error(t, err)
		assert.Nil(t, retrievedPayment)
		mockPaymentRepo.AssertExpectations(t)
	})
}
