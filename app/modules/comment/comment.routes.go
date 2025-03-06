package comment

import (
	"database/sql"
	"wordora/app/middlewares"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {
	commentRepo := NewCommentRepository(db)
	commentService := NewCommentService(commentRepo)
	commentController := NewCommentController(commentService)
	
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(tokenHelper))
	protected.POST("/", commentController.CreateComment)
	protected.PUT("/:id", commentController.UpdateComment)
	protected.DELETE("/:id", commentController.DeleteComment)
}
