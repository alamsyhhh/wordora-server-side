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

func (ctrl *CommentController) CreateComment(ctx *gin.Context) {
	ctrl.service.CreateComment(ctx)
}

func (ctrl *CommentController) UpdateComment(ctx *gin.Context) {
	ctrl.service.UpdateComment(ctx)
}

func (ctrl *CommentController) DeleteComment(ctx *gin.Context) {
	ctrl.service.DeleteComment(ctx)
}
