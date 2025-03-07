package category

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"wordora/app/modules/category/dto"
	"wordora/app/modules/category/model"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(ctx *gin.Context)
	GetAllCategories(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoryServiceImpl struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryServiceImpl{repo: repo}
}

func (s *categoryServiceImpl) GetAllCategories(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	name := ctx.Query("name")

	categories, total, err := s.repo.GetAllCategories(limit, offset, name)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch categories", err.Error())
		return
	}

	paginationResponse := common.GeneratePaginationResponse(categories, total, limit, offset)
	common.GenerateSuccessResponseWithData(ctx, "Categories retrieved successfully", paginationResponse)
}

func (s *categoryServiceImpl) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	category := model.Category{
		ID:         uuid.New().String(),
		Name:       req.Name,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	if err := s.repo.CreateCategory(&category); err != nil {
		log.Printf("Error inserting category: %v", err)
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to create category", nil)
		return
	}

	response := dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.Created_at,
		UpdatedAt: category.Updated_at,
	}

	common.GenerateSuccessResponseWithData(ctx, "Category created successfully", response)
}

func (s *categoryServiceImpl) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Category not found", nil)
		return
	}

	category.Name = req.Name
	category.Updated_at = time.Now()

	if err := s.repo.UpdateCategory(id, req.Name); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to update category", nil)
		return
	}

	response := dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.Created_at,
		UpdatedAt: category.Updated_at,
	}

	common.GenerateSuccessResponseWithData(ctx, "Category updated successfully", response)
}

func (s *categoryServiceImpl) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("Deleting category with ID:", id)

	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Category not found", nil)
		return
	}

	if err := s.repo.DeleteCategory(id); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete category", nil)
		return
	}

	response := dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.Created_at,
		UpdatedAt: category.Updated_at,
	}

	common.GenerateSuccessResponseWithData(ctx, "Category deleted successfully", response)
}

