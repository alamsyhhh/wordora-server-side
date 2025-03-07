package category

import (
	"database/sql"
	"errors"
	"log"
	"wordora/app/modules/category/model"

	"github.com/doug-martin/goqu/v9"
)

type CategoryRepository interface {
	CreateCategory(category *model.Category) error
	GetAllCategories(limit, offset int, search string) ([]model.Category, int, error)
	GetCategoryByID(id string) (*model.Category, error)
	UpdateCategory(id string, name string) error
	DeleteCategory(id string) error
}

type categoryRepositoryImpl struct {
	db *goqu.Database
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{db: goqu.New("postgres", db)}
}

func (r *categoryRepositoryImpl) CreateCategory(category *model.Category) error {
	_, err := r.db.Insert("categories").Rows(category).Executor().Exec()
	return err
}

func (r *categoryRepositoryImpl) GetAllCategories(limit, offset int, search string) ([]model.Category, int, error) {
	var categories []model.Category
	var err error
	query := r.db.From("categories")

	if search != "" {
		query = query.Where(goqu.Ex{"name": goqu.Op{"ilike": "%" + search + "%"}})
	}

	totalQuery := query.Select(goqu.COUNT("*"))
	var total int
	if _, err = totalQuery.ScanVal(&total); err != nil { 
		log.Println("Error counting articles:", err)
		return nil, 0, err
	}

	query = query.Limit(uint(limit)).Offset(uint(offset))

	err = query.ScanStructs(&categories)
	if err != nil {
		log.Println("Error fetching categories:", err)
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *categoryRepositoryImpl) GetCategoryByID(id string) (*model.Category, error) {
	var category model.Category
	found, err := r.db.From("categories").Where(goqu.Ex{"id": id}).ScanStruct(&category)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("category not found")
	}
	return &category, nil
}

func (r *categoryRepositoryImpl) UpdateCategory(id string, name string) error {
	_, err := r.db.Update("categories").
		Set(goqu.Record{"name": name, "updated_at": goqu.L("CURRENT_TIMESTAMP")}).
		Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}

func (r *categoryRepositoryImpl) DeleteCategory(id string) error {
	log.Println("Executing DELETE query for category ID:", id)
	_, err := r.db.Delete("categories").Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
