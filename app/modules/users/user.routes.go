package users

import (
	"database/sql"
	"wordora/app/middlewares"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userController := NewUserController(userService)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(tokenHelper))
	protected.GET("/me", userController.GetMe)
	protected.PUT("/:id/role", userController.UpdateUserRole)
}
