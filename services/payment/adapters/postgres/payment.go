package postgres

import (
	"context"
	"database/sql"

	"hexabank/internal/errors"
	"hexabank/services/payment/domain/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PaymentRepo struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) CreatePayment(ctx context.Context, payment *model.Payment) error {
	query := `INSERT INTO payments (id, description, amount, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, payment.ID, payment.Description, payment.Amount, payment.CreatedAt)
	return err
}

func (r *PaymentRepo) GetPayment(ctx context.Context, id uuid.UUID) (*model.Payment, error) {
	var payment model.Payment
	query := `SELECT id, description, amount, created_at FROM payments WHERE id = $1`
	err := r.db.GetContext(ctx, &payment, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundError
		}
		return nil, errors.InternalError
	}
	return &payment, err
}
