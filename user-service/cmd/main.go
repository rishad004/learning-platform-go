package main

import (
	"log"
	"net"

	grpchandlers "github.com/rishad004/learning-platform-go/user-service/internal/delivery/grpc-handlers"
	"github.com/rishad004/learning-platform-go/user-service/internal/repository"
	"github.com/rishad004/learning-platform-go/user-service/internal/service"
	"github.com/rishad004/learning-platform-go/user-service/pkg"
	pb "github.com/rishad004/learning-platform-go/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	Db := pkg.InitDatabase()

	repo := repository.NewUserRepository(Db)
	userService := service.NewUserService(repo)
	userHandler := grpchandlers.NewUserHandler(userService)

	g := grpc.NewServer()
	pb.RegisterUserServiceServer(g, userHandler)

	reflection.Register(g)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen, err: ", err)
	}

	log.Println("Server listening on port :8081")
	if err := g.Serve(listen); err != nil {
		log.Fatal("failed to listen - err: ", err)
	}
}
