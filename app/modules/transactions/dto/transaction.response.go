package dto

type SubscriptionResponse struct {
	ID        string `json:"id"`
	PlanID    string `json:"plan_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	IsActive  bool   `json:"is_active"`
}
