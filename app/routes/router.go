package routes

import (
	"database/sql"
	"time"
	"wordora/app/modules/article"
	"wordora/app/modules/auth"
	"wordora/app/modules/category"
	"wordora/app/modules/comment"
	"wordora/app/modules/reactions"
	"wordora/app/modules/users"

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
	article.SetupArticleRoutes(api.Group("/articles"), db, tokenHelper)
	reactions.SetupReactionsRoutes(api.Group("/reactions"), db, tokenHelper)
	comment.SetupCommentRoutes(api.Group("/comments"), db, tokenHelper)
	users.SetupUserRoutes(api.Group("/users"), db, tokenHelper)

	return router
}
