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

// CreateCategory godoc
// @Summary Create a new category
// @Description Add a new category to the system
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body dto.CreateCategoryRequest true "Create Category Request"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories [post]
// @Security BearerAuth
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	c.categoryService.CreateCategory(ctx)
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Param limit query int false "Number of articles per page" default(10)
// @Param offset query int false "Number of articles to skip" default(0)
// @Param name query string false "Search articles by name"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories [get]
// @Security BearerAuth
func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	c.categoryService.GetAllCategories(ctx)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param request body dto.UpdateCategoryRequest true "Update Category Request"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories/{id} [put]
// @Security BearerAuth
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	c.categoryService.UpdateCategory(ctx)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Remove a category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories/{id} [delete]
// @Security BearerAuth
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	c.categoryService.DeleteCategory(ctx)
}
