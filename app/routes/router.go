package routes

import (
	"database/sql"
	"time"
	"wordora/app/modules/auth"

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

	auth.SetupAuthRoutes(api.Group("/auth"), db)

	return router
}
