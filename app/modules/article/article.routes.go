package article

import (
	"database/sql"
	"wordora/app/middlewares"
	"wordora/app/modules/users"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {

	articleRepo := NewArticleRepository(db)
	userRepo := users.NewUserRepository(db)
	articleService := NewArticleService(articleRepo, userRepo)
	articleController := NewArticleController(articleService)

	router.GET("/", articleController.GetAllArticles)
	router.GET("/category/:category_id", articleController.GetArticlesByCategory)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(tokenHelper))
	protected.GET("/:id", articleController.GetArticleByID)

	Adminprotected := router.Group("/")
	Adminprotected.Use(middlewares.AdminOnly(tokenHelper))
	Adminprotected.POST("/", articleController.CreateArticle)
	Adminprotected.PUT("/:id", articleController.UpdateArticle)
	Adminprotected.DELETE("/:id", articleController.DeleteArticle)
}

