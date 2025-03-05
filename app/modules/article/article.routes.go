package article

import (
	"database/sql"
	"wordora/app/middlewares"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {

	articleRepo := NewArticleRepository(db)
	articleService := NewArticleService(articleRepo)
	articleController := NewArticleController(articleService)

	router.GET("/", articleController.GetAllArticles)
	router.GET("/:id", articleController.GetArticleByID)
	router.GET("/category/:category_id", articleController.GetArticlesByCategory)

	protected := router.Group("/")
	protected.Use(middlewares.AdminOnly(tokenHelper))
	router.POST("/", articleController.CreateArticle)
	router.PUT("/:id", articleController.UpdateArticle)
	router.DELETE("/:id", articleController.DeleteArticle)
}

