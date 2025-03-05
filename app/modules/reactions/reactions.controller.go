package reactions

import (
	"github.com/gin-gonic/gin"
)

type ReactionController struct {
	reactionService *ReactionService
}

func NewReactionController(reactionService *ReactionService) *ReactionController { 
	return &ReactionController{reactionService: reactionService}
}


func (c *ReactionController) CreateReact(ctx *gin.Context) {
	c.reactionService.CreateReact(ctx)
}

func (c *ReactionController) DeleteReact(ctx *gin.Context) {
	c.reactionService.DeleteReact(ctx)
}
