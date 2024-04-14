package main

import (
	"database/sql"
	"net"

	"github.com/aleroxac/goexpert-grpc/internal/database"
	"github.com/aleroxac/goexpert-grpc/internal/pb"
	"github.com/aleroxac/goexpert-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpServer, categoryService)
	reflection.Register(grpServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpServer.Serve(listener); err != nil {
		panic(err)
	}
}
