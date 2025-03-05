package transactions

import (
	"time"
	"wordora/app/modules/transactions/dto"
	"wordora/app/utils/midtrans"
	"wordora/app/utils/uuid"
)

type SubscriptionService interface {
	PurchaseSubscription(userID string, req dto.CreateSubscriptionRequest) (*Subscription, error)
	HandleMidtransWebhook(transactionID string, status string) error
}

type subscriptionService struct {
    repo     SubscriptionRepository
    midtrans *midtrans.MidtransService
}

func NewSubscriptionService(repo SubscriptionRepository, midtrans *midtrans.MidtransService) SubscriptionService {
    return &subscriptionService{repo, midtrans}
}


func (s *subscriptionService) PurchaseSubscription(userID string, req dto.CreateSubscriptionRequest) (*Subscription, error) {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 1, 0)

	sub := &Subscription{
		ID:        uuid.GenerateUUID(),
		UserID:    userID,
		PlanID:    req.PlanID,
		StartDate: startDate,
		EndDate:   endDate,
		IsActive:  true,
		CreatedAt: startDate,
		UpdatedAt: startDate,
	}

	err := s.repo.CreateSubscription(sub)
	if err != nil {
		return nil, err
	}

	token, _, err := s.midtrans.CreateTransaction(userID, sub.ID, 100000)
	if err != nil {
		return nil, err
	}

	tx := &Transaction{
		ID:                   uuid.GenerateUUID(),
		UserID:               userID,
		SubscriptionID:       sub.ID,
		MidtransTransactionID: token,
		Status:              "pending",
		GrossAmount:         100000,
		TransactionTime:     time.Now(),
	}

	err = s.repo.SaveTransaction(tx)
	return sub, err
}


func (s *subscriptionService) HandleMidtransWebhook(transactionID, status string) error {
	return s.repo.UpdateTransactionStatus(transactionID, status)
}
