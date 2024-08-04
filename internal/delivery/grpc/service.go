package grpc

import (
	"category-service/internal/model"
	pb "category-service/pb/category"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categoryUsecase model.CategoryUsecase
}

func NewCategoryService(categoryUsecase model.CategoryUsecase) *CategoryService {
	return &CategoryService{
		categoryUsecase: categoryUsecase,
	}
}
