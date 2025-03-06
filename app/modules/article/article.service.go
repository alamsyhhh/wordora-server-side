package article

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"wordora/app/modules/article/dto"
	"wordora/app/modules/article/model"
	"wordora/app/utils/cloudinary"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
)

type ArticleService interface {
	CreateArticle(ctx *gin.Context)
	GetAllArticles(ctx *gin.Context )
	GetArticleByID(ctx *gin.Context)
	UpdateArticle(ctx *gin.Context)
	DeleteArticle(ctx *gin.Context)
	GetArticlesByCategory(ctx *gin.Context)
}

type articleService struct {
	articleRepo ArticleRepository
}

func NewArticleService(articleRepo ArticleRepository) ArticleService {
	return &articleService{articleRepo: articleRepo}
}

func (s *articleService) CreateArticle(ctx *gin.Context) {
	var req dto.CreateArticleRequest
	if err := ctx.ShouldBind(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	var imagePath string
	file, _, err := ctx.Request.FormFile("image")
	if err == nil {
		imageBytes, err := io.ReadAll(file)
		if err != nil {
			common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to read image file", err.Error())
			return
		}

		uploadedURL, err := cloudinary.UploadImageToCloudinary(imageBytes)
		if err != nil {
			common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to upload image", err.Error())
			return
		}
		imagePath = uploadedURL
	}

	article := &model.Article{
		Title:      req.Title,
		CategoryID: req.CategoryID,
		Body:       req.Body,
		ImagePath:  imagePath,
	}

	savedArticle, err := s.articleRepo.CreateArticle(article)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to create article", err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Article created successfully", dto.ArticleResponse{
		ID:         savedArticle.ID,
		Title:      savedArticle.Title,
		CategoryID: savedArticle.CategoryID,
		Body:       savedArticle.Body,
		ImagePath:  savedArticle.ImagePath,
	})
}

func (s *articleService) GetAllArticles(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	name := ctx.Query("name")

	articles, total, err := s.articleRepo.GetAllArticles(limit, offset, name)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch articles", err.Error())
		return
	}

	paginationResponse := common.GeneratePaginationResponse(articles, total, limit, offset)
	common.GenerateSuccessResponseWithData(ctx, "Articles retrieved successfully", paginationResponse)
}

// func (s *articleService) GetAllArticles(ctx *gin.Context) {
// 	articles, err := s.articleRepo.GetAllArticles()
// 	if err != nil {
// 		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch articles", err.Error())
// 		return
// 	}
// 	common.GenerateSuccessResponseWithData(ctx, "Articles retrieved successfully", articles)
// }

// func (s *articleService) GetArticleByID(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	article, err := s.articleRepo.GetArticleByID(id)
// 	if err != nil {
// 		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Article not found", err.Error())
// 		return
// 	}
// 	common.GenerateSuccessResponseWithData(ctx, "Article retrieved successfully", article)
// }

func (s *articleService) GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")

	userID, exists := ctx.Get("sub")
	if !exists {

		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "User ID not found in token", nil)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Invalid user ID format", nil)
		return
	}

	articleDetail, err := s.articleRepo.GetArticleByIDWithDetails(id, userIDStr)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Article not found", err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Article retrieved successfully", articleDetail)
}

func (s *articleService) UpdateArticle(ctx *gin.Context) {
    id := ctx.Param("id")

    var req dto.UpdateArticleRequest
    if err := ctx.ShouldBind(&req); err != nil {
        common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err.Error())
        return
    }

    if req.CategoryID == "" {
        common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid input", "Category ID is required")
        return
    }

	var imagePath string
	file, _, err := ctx.Request.FormFile("image")
	if err == nil {
		imageBytes, err := io.ReadAll(file)
		if err != nil {
			common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to read image file", err.Error())
			return
		}

		uploadedURL, err := cloudinary.UploadImageToCloudinary(imageBytes)
		if err != nil {
			common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to upload image", err.Error())
			return
		}
		imagePath = uploadedURL
	}

    updatedArticle := &model.Article{
        Title:      req.Title,
        CategoryID: req.CategoryID,
        Body:       req.Body,
		ImagePath:  imagePath,
    }

    article, err := s.articleRepo.UpdateArticle(id, updatedArticle)
    if err != nil {
        fmt.Println("Update Error:", err)
        common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to update article", err.Error())
        return
    }

    common.GenerateSuccessResponseWithData(ctx, "Article updated successfully", article)
}

func (s *articleService) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	article, err := s.articleRepo.GetDeleteArticleByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Article not found", err.Error())
		return
	}

	err = s.articleRepo.DeleteArticle(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete article", err.Error())
		return
	}

	response := dto.ArticleResponse{
		ID:         article.ID,
		Title:      article.Title,
		CategoryID: article.CategoryID,
		Body:       article.Body,
		ImagePath:  article.ImagePath,
	}

	common.GenerateSuccessResponseWithData(ctx, "Article deleted successfully", response)
}


func (s *articleService) GetArticlesByCategory(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")

	articles, err := s.articleRepo.GetArticlesByCategory(categoryID)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch articles by category", err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "Articles by category retrieved successfully", articles)
}