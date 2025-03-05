package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReactionResponse struct {
	ID        uuid.UUID `json:"id"`
	ArticleID uuid.UUID `json:"article_id"`
	UserID    uuid.UUID `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
