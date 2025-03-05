package reactions

import (
	"net/http"
	"wordora/app/modules/reactions/dto"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReactionService struct {
	reactionRepo *ReactionRepository
}

func NewReactionService(reactionRepo *ReactionRepository) *ReactionService {
	return &ReactionService{reactionRepo: reactionRepo}
}

func (s *ReactionService) CreateReact(ctx *gin.Context) {
	userID, exists := ctx.Get("sub")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req dto.ReactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	existingReaction, _ := s.reactionRepo.GetReaction(req.ArticleID, uuid.MustParse(userID.(string)))
	if existingReaction != nil {
		common.GenerateErrorResponse(ctx, http.StatusConflict, "Reaction already exists", nil)
		return
	}

	reaction := &Reaction{
		ID:        uuid.New(),
		ArticleID: req.ArticleID,
		UserID:    uuid.MustParse(userID.(string)),
		Type:      req.Type,
	}

	err := s.reactionRepo.CreateReaction(reaction)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to add reaction", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Reaction added successfully", reaction)
}

func (s *ReactionService) DeleteReact(ctx *gin.Context) {
	userID, exists := ctx.Get("sub")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	articleID, err := uuid.Parse(ctx.Param("article_id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid article ID", nil)
		return
	}

	err = s.reactionRepo.DeleteReaction(articleID, uuid.MustParse(userID.(string)))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to remove reaction", nil)
		return
	}

	common.GenerateSuccessResponse(ctx, "Reaction removed successfully")
}
