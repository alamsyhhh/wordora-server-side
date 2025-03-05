package comment

import (
	"net/http"
	"time"
	"wordora/app/modules/comment/dto"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentService struct {
	repo *CommentRepository
}

func NewCommentService(repo *CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(ctx *gin.Context) {
	var req dto.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	userID, _ := ctx.Get("sub")
	comment := Comment{
		ID:        uuid.NewString(),
		ArticleID: req.ArticleID,
		UserID:    userID.(string),
		ParentID:  req.ParentID,
		Body:      req.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.CreateComment(&comment)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Comment created successfully", comment)
}

func (s *CommentService) UpdateComment(ctx *gin.Context) {
	var req dto.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	userID, _ := ctx.Get("sub")
	id := ctx.Param("id")

	comment, err := s.repo.GetCommentByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Comment not found", nil)
		return
	}

	if comment.UserID != userID.(string) {
		common.GenerateErrorResponse(ctx, http.StatusForbidden, "Unauthorized to edit this comment", nil)
		return
	}

	err = s.repo.UpdateComment(id, req.Body)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponse(ctx, "Comment updated successfully")
}

func (s *CommentService) DeleteComment(ctx *gin.Context) {
	userID, _ := ctx.Get("sub")
	id := ctx.Param("id")

	comment, err := s.repo.GetCommentByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Comment not found", nil)
		return
	}

	if comment.UserID != userID.(string) {
		common.GenerateErrorResponse(ctx, http.StatusForbidden, "Unauthorized to delete this comment", nil)
		return
	}

	err = s.repo.DeleteComment(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponse(ctx, "Comment deleted successfully")
}
