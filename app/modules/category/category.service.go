package category

import (
	"log"
	"time"
	"wordora/app/modules/category/dto"

	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetAllCategories() ([]dto.CategoryResponse, error)
	// GetCategoryByID(id string) (*dto.CategoryResponse, error)
	UpdateCategory(id string, req dto.UpdateCategoryRequest) error
	DeleteCategory(id string) error
}

type categoryServiceImpl struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryServiceImpl{repo: repo}
}

func (s *categoryServiceImpl) CreateCategory(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := Category{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	if err := s.repo.CreateCategory(&category); err != nil {
		log.Printf("Error inserting category: %v", err)
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:      category.ID,
		Name:    category.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *categoryServiceImpl) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	var responses []dto.CategoryResponse
	for _, category := range categories {
		responses = append(responses, dto.CategoryResponse{
			ID:      category.ID,
			Name:    category.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return responses, nil
}

// func (s *categoryServiceImpl) GetCategoryByID(id string) (*dto.CategoryResponse, error) {
// 	category, err := s.repo.GetCategoryByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dto.CategoryResponse{
// 		ID:      category.ID,
// 		Name:    category.Name,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}, nil
// }

func (s *categoryServiceImpl) UpdateCategory(id string, req dto.UpdateCategoryRequest) error {
	return s.repo.UpdateCategory(id, req.Name)
}

func (s *categoryServiceImpl) DeleteCategory(id string) error {
	return s.repo.DeleteCategory(id)
}
