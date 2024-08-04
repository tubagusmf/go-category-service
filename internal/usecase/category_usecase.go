package usecase

import "category-service/internal/model"

type categoryUsecase struct {
	categoryRepo model.CategoryRepository
}

func NewCategoryUsecase(cr model.CategoryRepository) model.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: cr,
	}
}

func (cu *categoryUsecase) FindCategoryById(id string) (model.Category, error) {
	return cu.categoryRepo.FindCategoryById(id)
}
