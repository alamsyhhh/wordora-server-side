package auth

import (
	"database/sql"
	"wordora/app/modules/profiles"
	"wordora/app/modules/users"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup, db *sql.DB) {
	userRepo := users.NewUserRepository(db)
	profileRepo := profiles.NewProfileRepository(db)

	authService := NewAuthService(userRepo, profileRepo, paseto.NewTokenHelper())
	authController := NewAuthController(authService)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
}
