package routes

import (
	"database/sql"
	"time"
	"wordora/app/modules/auth"
	"wordora/app/modules/category"

	"wordora/app/utils/paseto"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:          12 * time.Hour,
	}))

	api := router.Group("/v1/api")

	tokenHelper := paseto.NewTokenHelper()

	auth.SetupAuthRoutes(api.Group("/auth"), db)
	category.SetupCategoryRoutes(api.Group("/categories"), db, tokenHelper)

	return router
}
