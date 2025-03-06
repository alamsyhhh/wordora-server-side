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

// CreateReact godoc
// @Summary Add a reaction
// @Description Add a reaction (like, love, etc.) to a post or comment
// @Tags Reactions
// @Accept json
// @Produce json
// @Param request body dto.ReactionRequest true "Reaction Request"
// @Success 200 {object} map[string]interface{} "Reaction added successfully"
// @Router /reactions [post]
// @Security BearerAuth
func (c *ReactionController) CreateReact(ctx *gin.Context) {
	c.reactionService.CreateReact(ctx)
}

// DeleteReact godoc
// @Summary Remove a reaction
// @Description Remove a reaction from a post or comment
// @Tags Reactions
// @Accept json
// @Produce json
// @Param id path string true "Reaction ID"
// @Success 200 {object} map[string]interface{} "Reaction removed successfully"
// @Router /reactions/{id} [delete]
// @Security BearerAuth
func (c *ReactionController) DeleteReact(ctx *gin.Context) {
	c.reactionService.DeleteReact(ctx)
}
