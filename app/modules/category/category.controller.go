package category

import (
	"net/http"
	"wordora/app/modules/category/dto"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	category, err := c.service.CreateCategory(req)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to create category", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Category created successfully", category)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch categories", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Categories fetched successfully", categories)
}

// func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	category, err := c.service.GetCategoryByID(id)
// 	if err != nil {
// 		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Category not found", nil)
// 		return
// 	}

// 	common.GenerateSuccessResponseWithData(ctx, "Category found", category)
// }

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.service.UpdateCategory(id, req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to update category", nil)
		return
	}

	common.GenerateSuccessResponse(ctx, "Category updated successfully")
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteCategory(id); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete category", nil)
		return
	}

	common.GenerateSuccessResponse(ctx, "Category deleted successfully")
}
