package transactions

import (
	"time"
)

type Subscription struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	PlanID    string    `db:"plan_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Transaction struct {
	ID                   string    `db:"id"`
	UserID               string    `db:"user_id"`
	SubscriptionID       string    `db:"subscription_id"`
	MidtransTransactionID string    `db:"midtrans_transaction_id"`
	Status              string    `db:"status"`
	PaymentType         string    `db:"payment_type"`
	GrossAmount         int       `db:"gross_amount"`
	TransactionTime     time.Time `db:"transaction_time"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

