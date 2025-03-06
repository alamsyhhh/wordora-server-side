package article

import (
	"database/sql"
	"log"
	"time"
	"wordora/app/modules/article/dto"
	"wordora/app/modules/article/model"
	"wordora/app/modules/comment"
	"wordora/app/modules/reactions"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type ArticleRepository interface {
	CreateArticle(article *model.Article) (*model.Article, error)
	GetAllArticles(limit, offset int, search string) ([]model.Article, int, error)
	// GetArticleByID(id string, userID string) (*ArticleDetailResponse, error)
	GetDeleteArticleByID(id string) (*model.Article, error)
	GetArticleByIDWithDetails(articleID, userID string) (*dto.ArticleDetailResponse, error)
	UpdateArticle(id string, updatedArticle *model.Article) (*model.Article, error)
	DeleteArticle(id string) error
	GetArticlesByCategory(categoryID string) ([]model.Article, error)
}

type articleRepository struct {
	db *goqu.Database
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{
		db: goqu.New("postgres", db),
	}
}

func (r *articleRepository) CreateArticle(article *model.Article) (*model.Article, error) {
	article.ID = uuid.NewString()

	query, args, err := goqu.Insert("articles").
		Rows(article).
		Returning("id", "title", "category_id", "body", "image_path").
		ToSQL()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	var newArticle model.Article
	err = r.db.QueryRow(query, args...).Scan(&newArticle.ID, &newArticle.Title, &newArticle.CategoryID, &newArticle.Body, &newArticle.ImagePath)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return &newArticle, nil
}

func (r *articleRepository) GetAllArticles(limit, offset int, search string) ([]model.Article, int, error) {
	var articles []model.Article
	var err error
	query := r.db.From("articles")

	if search != "" {
		query = query.Where(goqu.Ex{"title": goqu.Op{"ilike": "%" + search + "%"}})
	}

	totalQuery := query.Select(goqu.COUNT("*"))
	var total int
	if _, err = totalQuery.ScanVal(&total); err != nil { 
		log.Println("Error counting articles:", err)
		return nil, 0, err
	}

	query = query.Limit(uint(limit)).Offset(uint(offset))

	err = query.ScanStructs(&articles)
	if err != nil {
		log.Println("Error fetching articles:", err)
		return nil, 0, err
	}

	return articles, total, nil
}


// func (r *articleRepository) GetAllArticles() ([]model.Article, error) {
// 	var articles []model.Article
// 	err := r.db.From("articles").ScanStructs(&articles)
// 	if err != nil {
// 		log.Println("Error fetching articles:", err)
// 		return nil, err
// 	}
// 	return articles, nil
// }

func (r *articleRepository) GetDeleteArticleByID(id string) (*model.Article, error) {
	var article model.Article
	found, err := r.db.From("articles").Where(goqu.Ex{"id": id}).ScanStruct(&article)
	if err != nil {
		log.Println("Error fetching article:", err)
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &article, nil
}

func (r *articleRepository) GetArticleByIDWithDetails(articleID, userID string) (*dto.ArticleDetailResponse, error) {
	var article model.Article
	found, err := r.db.From("articles").
		Select("id", "title", "category_id", "body", "image_path", "created_at", "updated_at").
		Where(goqu.Ex{"id": articleID}).
		ScanStruct(&article)
	if err != nil {
		log.Println("Error fetching article:", err)
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}

	// Ambil semua komentar untuk artikel ini
	var comments []comment.Comment
	err = r.db.From("comments").
		Select("id", "article_id", "user_id", "parent_id", "body", "created_at", "updated_at").
		Where(goqu.Ex{"article_id": articleID}).
		ScanStructs(&comments)
	if err != nil {
		log.Println("Error fetching comments:", err)
		return nil, err
	}

	// Ambil semua reaksi dari user yang login
	var reactions []reactions.Reaction
	err = r.db.From("reactions").
		Select("id", "article_id", "user_id", "type", "created_at", "updated_at").
		Where(goqu.Ex{"article_id": articleID, "user_id": userID}).
		ScanStructs(&reactions)
	if err != nil {
		log.Println("Error fetching reactions:", err)
		return nil, err
	}

	// Gabungkan data ke dalam response
	response := &dto.ArticleDetailResponse{
		Article:   article,
		Comments:  comments,
		Reactions: reactions,
	}

	return response, nil
}


func (r *articleRepository) UpdateArticle(id string, updatedArticle *model.Article) (*model.Article, error) {
	query, args, err := goqu.Update("articles").
		Set(goqu.Record{
			"title":       updatedArticle.Title,
			"category_id": updatedArticle.CategoryID,
			"body":        updatedArticle.Body,
			"image_path":  updatedArticle.ImagePath,
			"updated_at":  time.Now(),
		}).
		Where(goqu.Ex{"id": id}).
		Returning("*").
		ToSQL()
	if err != nil {
		return nil, err
	}

	var article model.Article
	err = r.db.QueryRow(query, args...).Scan(
		&article.ID,
		&article.Title,
		&article.CategoryID,
		&article.Body,
		&article.ImagePath,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) DeleteArticle(id string) error {
	_, err := r.db.Delete("articles").Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}

func (r *articleRepository) GetArticlesByCategory(categoryID string) ([]model.Article, error) {
	var articles []model.Article

	log.Println("Executing query with category_id:", categoryID)
	err := r.db.From("articles").Where(goqu.Ex{"category_id": categoryID}).ScanStructs(&articles)
	if err != nil {
		log.Println("Error fetching articles by category:", err)
		return nil, err
	}
	return articles, nil
}