package article

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type ArticleRepository interface {
	CreateArticle(article *Article) (*Article, error)
	GetAllArticles() ([]Article, error)
	GetArticleByID(id string) (*Article, error)
	UpdateArticle(id string, updatedArticle *Article) (*Article, error)
	DeleteArticle(id string) error
	GetArticlesByCategory(categoryID string) ([]Article, error)
}

type articleRepository struct {
	db *goqu.Database
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{
		db: goqu.New("postgres", db),
	}
}

func (r *articleRepository) CreateArticle(article *Article) (*Article, error) {
	article.ID = uuid.NewString()

	// Generate SQL query
	query, args, err := goqu.Insert("articles").
		Rows(article).
		Returning("id", "title", "category_id", "body", "image_path").
		ToSQL()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	// Eksekusi query menggunakan `sql.DB`
	var newArticle Article
	err = r.db.QueryRow(query, args...).Scan(&newArticle.ID, &newArticle.Title, &newArticle.CategoryID, &newArticle.Body, &newArticle.ImagePath)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return &newArticle, nil
}

func (r *articleRepository) GetAllArticles() ([]Article, error) {
	var articles []Article
	err := r.db.From("articles").ScanStructs(&articles)
	if err != nil {
		log.Println("Error fetching articles:", err)
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetArticleByID(id string) (*Article, error) {
	var article Article
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

func (r *articleRepository) UpdateArticle(id string, updatedArticle *Article) (*Article, error) {
	query, args, err := goqu.Update("articles").
		Set(goqu.Record{
			"title":       updatedArticle.Title,
			"category_id": updatedArticle.CategoryID,
			"body":        updatedArticle.Body,
			"image_path":  updatedArticle.ImagePath,
		}).
		Where(goqu.Ex{"id": id}).
		Returning("*").
		ToSQL()
	if err != nil {
		return nil, err
	}

	var article Article
	err = r.db.QueryRow(query, args...).Scan(&article.ID, &article.Title, &article.CategoryID, &article.Body, &article.ImagePath)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) DeleteArticle(id string) error {
	_, err := r.db.Delete("articles").Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}

func (r *articleRepository) GetArticlesByCategory(categoryID string) ([]Article, error) {
	var articles []Article
	err := r.db.From("articles").Where(goqu.Ex{"category_id": categoryID}).ScanStructs(&articles)
	if err != nil {
		log.Println("Error fetching articles by category:", err)
		return nil, err
	}
	return articles, nil
}