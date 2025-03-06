package article

import (
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleService ArticleService
}

func NewArticleController(articleService ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with title, content, category, and image
// @Tags Articles
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Article Title"
// @Param category_id formData string true "Category ID"
// @Param body formData string true "Article Body"
// @Param image formData file false "Article Image"
// @Success 200 {object} map[string]interface{} "Article created successfully"
// @Router /articles [post]
// @Security BearerAuth
func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	c.articleService.CreateArticle(ctx)
}

// GetAllArticles godoc
// @Summary Get all articles
// @Description Retrieve a list of all articles
// @Tags Articles
// @Produce json
// @Success 200 {array} map[string]interface{} "List of articles"
// @Router /articles [get]
func (c *ArticleController) GetAllArticles(ctx *gin.Context) {
	c.articleService.GetAllArticles(ctx)
}

// GetArticleByID godoc
// @Summary Get an article by ID
// @Description Retrieve a single article by its ID
// @Tags Articles
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} map[string]interface{} "Article retrieved successfully"
// @Router /articles/{id} [get]
func (c *ArticleController) GetArticleByID(ctx *gin.Context) {
	c.articleService.GetArticleByID(ctx)
}

// UpdateArticle godoc
// @Summary Update an article
// @Description Update article details like title, content, category, and image
// @Tags Articles
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Article ID"
// @Param title formData string true "Article Title"
// @Param category_id formData string true "Category ID"
// @Param body formData string true "Article Body"
// @Param image formData file false "Article Image"
// @Success 200 {object} map[string]interface{} "Article updated successfully"
// @Router /articles/{id} [put]
// @Security BearerAuth
func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	c.articleService.UpdateArticle(ctx)
}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article by its ID
// @Tags Articles
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} map[string]interface{} "Article deleted successfully"
// @Router /articles/{id} [delete]
// @Security BearerAuth
func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	c.articleService.DeleteArticle(ctx)
}

// GetArticlesByCategory godoc
// @Summary Get articles by category
// @Description Retrieve a list of articles filtered by category
// @Tags Articles
// @Produce json
// @Param category_id path string true "Category ID"
// @Success 200 {array} map[string]interface{} "List of articles in the category"
// @Router /articles/category/{category_id} [get]
func (c *ArticleController) GetArticlesByCategory(ctx *gin.Context) {
	c.articleService.GetArticlesByCategory(ctx)
}

