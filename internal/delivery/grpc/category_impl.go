package grpc

import (
	"context"
	"errors"

	pb "github.com/tubagusmf/go-category-service/pb/category"
)

// FindCategorys is a function to find categorys
func (s *CategoryService) FindCategoryById(ctx context.Context, in *pb.CategoryRequest) (*pb.Category, error) {
	if in.Id == "" {
		return nil, errors.New("id is required")
	}
	category, err := s.categoryUsecase.FindCategoryById(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
