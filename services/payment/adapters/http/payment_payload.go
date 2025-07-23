package http

type CreatePaymentPayload struct {
	Description string `json:"description" validate:"required"`
	Amount      int    `json:"amount" validate:"required,gt=0"`
}
