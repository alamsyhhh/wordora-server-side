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
	router.POST("/", commentController.CreateComment)
	router.PUT("/:id", commentController.UpdateComment)
	router.DELETE("/:id", commentController.DeleteComment)
}
