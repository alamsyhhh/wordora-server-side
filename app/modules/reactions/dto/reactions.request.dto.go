package dto

import "github.com/google/uuid"

type ReactionRequest struct {
	ArticleID uuid.UUID `json:"article_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Type      string    `json:"type" binding:"required,oneof=like love clap insightful funny sad" example:"like"`
}
