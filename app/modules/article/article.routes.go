package article

import (
	"database/sql"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {

	articleRepo := NewArticleRepository(db)
	articleService := NewArticleService(articleRepo)
	articleController := NewArticleController(articleService)

	router.POST("/", articleController.CreateArticle)
	router.GET("/", articleController.GetAllArticles)
	router.GET("/:id", articleController.GetArticleByID)
	router.PUT("/:id", articleController.UpdateArticle)
	router.DELETE("/:id", articleController.DeleteArticle)
	router.GET("/category/:category_id", articleController.GetArticlesByCategory)
}

