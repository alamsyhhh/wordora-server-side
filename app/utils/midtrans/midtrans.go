package midtrans

import (
	"errors"
	"fmt"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService struct{}

func NewMidtransService() *MidtransService {
	return &MidtransService{}
}

func (ms *MidtransService) CreateTransaction(userID, subscriptionID string, amount int) (string, string, error) {
	snapClient := snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("%s-%s", userID, subscriptionID),
			GrossAmt: int64(amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: userID, // Gunakan FName karena CustomerID tidak valid
		},
	}

	snapResp, err := snapClient.CreateTransaction(req)
	if err != nil {
		return "", "", err
	}

	return snapResp.Token, snapResp.RedirectURL, nil
}

func (ms *MidtransService) ProcessWebhook(notification map[string]interface{}) error {
	transactionStatus, exists := notification["transaction_status"].(string)
	if !exists {
		return errors.New("invalid notification format")
	}

	orderID, exists := notification["order_id"].(string)
	if !exists {
		return errors.New("invalid order ID")
	}

	// Pastikan variabel digunakan agar tidak error unused variable
	fmt.Println("Transaction Status:", transactionStatus)
	fmt.Println("Order ID:", orderID)

	// Implementasikan update transaksi di database

	return nil
}
