package transactions

import (
	"net/http"
	"wordora/app/modules/transactions/dto"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
)

type SubscriptionController struct {
	service SubscriptionService
}

func NewSubscriptionController(service SubscriptionService) *SubscriptionController {
	return &SubscriptionController{service}
}

func (sc *SubscriptionController) PurchaseSubscription(c *gin.Context) {
	userID := c.GetString("sub")

	var req dto.CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	sub, err := sc.service.PurchaseSubscription(userID, req)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, "Failed to create subscription", nil)
		return
	}

	common.GenerateSuccessResponseWithData(c, "Subscription created", sub)
}

func (sc *SubscriptionController) WebhookHandler(c *gin.Context) {
	var payload struct {
		TransactionID string `json:"transaction_id"`
		Status        string `json:"status"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, "Invalid webhook payload", nil)
		return
	}

	err := sc.service.HandleMidtransWebhook(payload.TransactionID, payload.Status)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, "Failed to update transaction", nil)
		return
	}

	common.GenerateSuccessResponse(c, "Transaction updated successfully")
}
