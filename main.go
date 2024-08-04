package main

import (
	"log"
	"net"

	"github.com/tubagusmf/go-category-service/internal/usecase"

	"github.com/tubagusmf/go-category-service/internal/repository"

	grpcSvc "github.com/tubagusmf/go-category-service/internal/delivery/grpc"
	pb "github.com/tubagusmf/go-category-service/pb/category"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	categoryRepo := repository.NewCategoryRepository()
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryService := grpcSvc.NewCategoryService(categoryUsecase)

	pb.RegisterCategoryServiceServer(s, categoryService)

	listen, err := net.Listen("tcp", ":3300")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on port 3300")

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
