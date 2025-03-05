package transactions

import (
	"github.com/doug-martin/goqu/v9"
)

type SubscriptionRepository interface {
	CreateSubscription(sub *Subscription) error
	SaveTransaction(tx *Transaction) error
	UpdateTransactionStatus(transactionID, status string) error
}

type subscriptionRepository struct {
	db *goqu.Database
}

func NewSubscriptionRepository(db *goqu.Database) SubscriptionRepository {
	return &subscriptionRepository{db}
}

func (r *subscriptionRepository) CreateSubscription(sub *Subscription) error {
	_, err := r.db.Insert("subscriptions").Rows(sub).Executor().Exec()
	return err
}

func (r *subscriptionRepository) SaveTransaction(tx *Transaction) error {
	_, err := r.db.Insert("transactions").Rows(tx).Executor().Exec()
	return err
}

func (r *subscriptionRepository) UpdateTransactionStatus(transactionID, status string) error {
	_, err := r.db.Update("transactions").Set(goqu.Record{"status": status}).Where(goqu.Ex{"midtrans_transaction_id": transactionID}).Executor().Exec()
	return err
}
