package http

import (
	"context"
	"net/http"

	"hexabank/internal/errors"
	"hexabank/services/payment/domain/port"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PaymentHTTP struct {
	paymentService port.PaymentService
	validate       *validator.Validate
}

func NewPaymentHTTP(paymentService port.PaymentService) *PaymentHTTP {
	return &PaymentHTTP{
		paymentService: paymentService,
		validate:       validator.New(),
	}
}

func (h *PaymentHTTP) RegisterRoutes(r *gin.Engine) {
	endpoint := r.Group("/api/v1")
	endpoint.POST("/payments", h.CreatePaymentHandler)
	endpoint.GET("/payments/:id", h.GetPaymentHandler)
}

func (h *PaymentHTTP) CreatePaymentHandler(c *gin.Context) {
	var payload CreatePaymentPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	description := payload.Description
	amount := payload.Amount

	paymentRecord, err := h.paymentService.CreatePayment(context.Background(), description, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fullfil payment", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully", "payment": paymentRecord})
}

func (h *PaymentHTTP) GetPaymentHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
		return
	}

	paymentRecord, err := h.paymentService.GetPayment(context.Background(), id)
	if err != nil {
		if err == errors.NotFoundError {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment retrieved successfully", "payment": paymentRecord})
}
