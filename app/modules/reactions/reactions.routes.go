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
	router.POST("/", reactionController.CreateReact)
	router.DELETE("/:article_id", reactionController.DeleteReact)
}
