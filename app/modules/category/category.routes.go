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

	// adminAuthorization := middlewares.AdminOnly(tokenHelper)
	// authMiddleware := middlewares.AuthMiddleware(tokenHelper)

	// router.POST("/", adminAuthorization, categoryController.CreateCategory)

	router.GET("/", categoryController.GetAllCategories)
	// protected := router.Group("/")
	// protected.Use(authMiddleware)
	// protected.GET("/", categoryController.GetAllCategories)

	// router.GET("/:id", categoryController.GetCategoryByID)
	// router.PUT("/:id", adminAuthorization, categoryController.UpdateCategory)
	// router.DELETE("/:id", adminAuthorization, categoryController.DeleteCategory)

	protected := router.Group("/")
	protected.Use(middlewares.AdminOnly(tokenHelper))
	protected.POST("/", categoryController.CreateCategory)
	protected.PUT("/:id", categoryController.CreateCategory)
	protected.DELETE("/:id", categoryController.CreateCategory)
}
