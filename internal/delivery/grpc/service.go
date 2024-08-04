package grpc

import (
	"github.com/tubagusmf/go-category-service/internal/model"
	pb "github.com/tubagusmf/go-category-service/pb/category"
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
