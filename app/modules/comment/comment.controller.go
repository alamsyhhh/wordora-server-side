package comment

import (
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service *CommentService
}

func NewCommentController(service *CommentService) *CommentController {
	return &CommentController{service: service}
}

// CreateComment godoc
// @Summary Create a new comment
// @Description Add a new comment to a post or article
// @Tags Comments
// @Accept json
// @Produce json
// @Param request body dto.CreateCommentRequest true "Comment Request"
// @Success 200 {object} map[string]interface{} "Comment created successfully"
// @Router /comments [post]
// @Security BearerAuth
func (ctrl *CommentController) CreateComment(ctx *gin.Context) {
	ctrl.service.CreateComment(ctx)
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update an existing comment by ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param request body dto.UpdateCommentRequest true "Comment Request"
// @Success 200 {object} map[string]interface{} "Comment updated successfully"
// @Router /comments/{id} [put]
// @Security BearerAuth
func (ctrl *CommentController) UpdateComment(ctx *gin.Context) {
	ctrl.service.UpdateComment(ctx)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Remove a comment by ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} map[string]interface{} "Comment deleted successfully"
// @Router /comments/{id} [delete]
// @Security BearerAuth
func (ctrl *CommentController) DeleteComment(ctx *gin.Context) {
	ctrl.service.DeleteComment(ctx)
}
