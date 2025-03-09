package category

import (
	"database/sql"
	"wordora/app/middlewares"
	"wordora/app/utils/paseto"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.RouterGroup, db *sql.DB, tokenHelper *paseto.TokenHelper) {
	categoryRepo := NewCategoryRepository(db)
	categoryService := NewCategoryService(categoryRepo)
	categoryController := NewCategoryController(categoryService)

	router.GET("/", categoryController.GetAllCategories)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(tokenHelper))
	protected.POST("/", categoryController.CreateCategory)
	protected.PUT("/:id", categoryController.UpdateCategory)
	protected.DELETE("/:id", categoryController.DeleteCategory)
}
