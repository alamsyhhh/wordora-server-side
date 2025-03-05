package category

import (
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService CategoryService
}

func NewCategoryController(categoryService CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	c.categoryService.CreateCategory(ctx)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	c.categoryService.GetAllCategories(ctx)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	c.categoryService.UpdateCategory(ctx)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	c.categoryService.DeleteCategory(ctx)
}
