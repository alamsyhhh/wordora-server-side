package reactions

import (
	"database/sql"
	"wordora/app/middlewares"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupReactionsRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {
	reactionRepo := NewReactionRepository(db)
	reactionService := NewReactionService(reactionRepo)
	reactionController := NewReactionController(reactionService)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(tokenHelper))
	protected.POST("/", reactionController.CreateReact)
	protected.DELETE("/:reaction_id", reactionController.DeleteReact)
}
