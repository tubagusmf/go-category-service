package main

import (
	"category-service/internal/repository"
	"category-service/internal/usecase"
	"log"
	"net"

	grpcSvc "category-service/internal/delivery/grpc"
	pb "category-service/pb/category"

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
