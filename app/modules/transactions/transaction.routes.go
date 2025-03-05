package transactions

import (
	"wordora/app/middlewares"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupSubscriptionRoutes(router *gin.Engine, sc *SubscriptionController, tokenHelper *paseto.TokenHelper) {
	sub := router.Group("/subscriptions")
	sub.Use(middlewares.AuthMiddleware(tokenHelper))
	{
		sub.POST("/purchase", sc.PurchaseSubscription)
		sub.POST("/webhook", sc.WebhookHandler)
	}
}
