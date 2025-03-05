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

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	c.articleService.CreateArticle(ctx)
}

func (c *ArticleController) GetAllArticles(ctx *gin.Context) {
	c.articleService.GetAllArticles(ctx)
}

func (c *ArticleController) GetArticleByID(ctx *gin.Context) {
	c.articleService.GetArticleByID(ctx)
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	c.articleService.UpdateArticle(ctx)
}

func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	c.articleService.DeleteArticle(ctx)
}

func (c *ArticleController) GetArticlesByCategory(ctx *gin.Context) {
	c.articleService.GetArticlesByCategory(ctx)
}