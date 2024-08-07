package main

import (
	"log"
	"net"

	grpchandlers "github.com/rishad004/learning-platform-go/institute-service/internal/handler/grpc-handlers"
	"github.com/rishad004/learning-platform-go/institute-service/internal/repository"
	"github.com/rishad004/learning-platform-go/institute-service/internal/service"
	"github.com/rishad004/learning-platform-go/institute-service/pkg"
	pb "github.com/rishad004/learning-platform-go/institute-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	Db := pkg.InitDatabase()

	repo := repository.NewInstituteRepository(Db)
	instituteService := service.NewInstituteService(repo)
	instituteHandler := grpchandlers.NewInstituteHandler(instituteService)

	g := grpc.NewServer()
	pb.RegisterAdminServiceServer(g, instituteHandler)

	reflection.Register(g)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("failed to listen, err: ", err)
	}

	log.Println("Server listening on port :8081")
	if err := g.Serve(listen); err != nil {
		log.Fatal("failed to listen - err: ", err)
	}
}
