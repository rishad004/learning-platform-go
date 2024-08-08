package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	grpchandlers "github.com/rishad004/learning-platform-go/user-service/internal/handler/grpc-handlers"
	httpcontrollers "github.com/rishad004/learning-platform-go/user-service/internal/handler/http-controllers"
	"github.com/rishad004/learning-platform-go/user-service/internal/repository"
	"github.com/rishad004/learning-platform-go/user-service/internal/service"
	"github.com/rishad004/learning-platform-go/user-service/pkg"
	pb "github.com/rishad004/learning-platform-go/user-service/proto"
	institute_pb "github.com/rishad004/learning-platform-go/user-service/proto/client/institute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	Db := pkg.InitDatabase()

	connHotel, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer connHotel.Close()
	hotelService := institute_pb.NewAdminServiceClient(connHotel)

	repo := repository.NewUserRepository(Db)
	userService := service.NewUserService(repo, hotelService)
	userHandler := grpchandlers.NewUserHandler(userService)
	httpController := httpcontrollers.NewUserController(userService)

	g := grpc.NewServer()
	pb.RegisterUserServiceServer(g, userHandler)

	reflection.Register(g)

	go func() {
		listen, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatal("failed to listen, err: ", err)
		}

		log.Println("Server listening on port :8081")
		if err := g.Serve(listen); err != nil {
			log.Fatal("failed to listen - err: ", err)
		}
	}()

	go func() {
		r := gin.Default()
		httpController.UserRouters(r)

		if err := r.Run(":8080"); err != nil {
			log.Fatal("failed to listen, err: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
