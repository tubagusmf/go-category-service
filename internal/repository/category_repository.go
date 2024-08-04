package repository

import "github.com/tubagusmf/category-service/internal/model"

type categoryRepository struct {
}

func NewCategoryRepository() model.CategoryRepository {
	return &categoryRepository{}
}

func (cr *categoryRepository) FindCategoryById(id string) (model.Category, error) {
	var category model.Category

	if id != "1" {
		return category, nil
	}

	return model.Category{
		Id:        "1",
		Name:      "Technology",
		CreatedAt: "2022-01-01T00:00:00Z",
		UpdatedAt: "2022-01-01T00:00:00Z",
	}, nil
}
