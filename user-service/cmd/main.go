package main

import (
	grpchandlers "github.com/rishad004/learning-platform-go/user-service/internal/delivery/grpc-handlers"
	"github.com/rishad004/learning-platform-go/user-service/internal/repository"
	"github.com/rishad004/learning-platform-go/user-service/internal/service"
	"github.com/rishad004/learning-platform-go/user-service/pkg"
	pb "github.com/rishad004/learning-platform-go/user-service/proto"
	"google.golang.org/grpc"
)

func main() {

	Db := pkg.InitDatabase()

	repo := repository.NewUserRepository(Db)
	userService := service.NewUserService(repo)
	userHandler := grpchandlers.NewUserHandler(userService)

	g := grpc.NewServer()
	pb.RegisterUserServiceServer(g, userHandler)
}
