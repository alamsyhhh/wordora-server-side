package dto

import (
	"wordora/app/modules/article/model"
	"wordora/app/modules/comment"
	"wordora/app/modules/reactions"
)

type ArticleDetailResponse struct {
	Article   model.Article    `json:"article"`
	Comments  []comment.Comment  `json:"comments"`
	Reactions []reactions.Reaction `json:"user_reactions"`
}