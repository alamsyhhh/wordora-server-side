package category

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type CategoryRepository interface {
	CreateCategory(category *Category) error
	GetAllCategories() ([]Category, error)
	// GetCategoryByID(id string) (*Category, error)
	UpdateCategory(id string, name string) error
	DeleteCategory(id string) error
}

type categoryRepositoryImpl struct {
	db *goqu.Database
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{db: goqu.New("postgres", db)}
}

func (r *categoryRepositoryImpl) CreateCategory(category *Category) error {
	_, err := r.db.Insert("categories").Rows(category).Executor().Exec()
	return err
}

func (r *categoryRepositoryImpl) GetAllCategories() ([]Category, error) {
	var categories []Category
	err := r.db.From("categories").ScanStructs(&categories)
	return categories, err
}

// func (r *categoryRepositoryImpl) GetCategoryByID(id string) (*Category, error) {
// 	var category Category
// 	found, err := r.db.From("categories").Where(goqu.Ex{"id": id}).ScanStruct(&category)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !found {
// 		return nil, errors.New("category not found")
// 	}
// 	return &category, nil
// }

func (r *categoryRepositoryImpl) UpdateCategory(id string, name string) error {
	_, err := r.db.Update("categories").
		Set(goqu.Record{"name": name, "updated_at": goqu.L("CURRENT_TIMESTAMP")}).
		Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}

func (r *categoryRepositoryImpl) DeleteCategory(id string) error {
	_, err := r.db.Delete("categories").Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
