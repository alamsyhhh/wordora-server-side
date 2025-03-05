package dto

type SubscriptionRequest struct {
	PlanID string `json:"plan_id" binding:"required"`
}
