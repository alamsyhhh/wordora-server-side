package dto

import "github.com/google/uuid"

type ReactionRequest struct {
	ArticleID uuid.UUID `json:"article_id" binding:"required"`
	Type      string    `json:"type" binding:"required,oneof=like love clap insightful funny sad"`
}
