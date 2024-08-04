package grpc

import (
	"github.com/tubagusmf/category-service/internal/model"
	pb "github.com/tubagusmf/category-service/pb/category"
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
