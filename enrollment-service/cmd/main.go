package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	grpchandlers "github.com/rishad004/learning-platform-go/enrollment-service/internal/handler/grpc-handlers"
	httpcontrollers "github.com/rishad004/learning-platform-go/enrollment-service/internal/handler/http-controllers"
	"github.com/rishad004/learning-platform-go/enrollment-service/internal/repository"
	"github.com/rishad004/learning-platform-go/enrollment-service/internal/service"
	"github.com/rishad004/learning-platform-go/enrollment-service/pkg"
	pb "github.com/rishad004/learning-platform-go/enrollment-service/proto"
	payment_pb "github.com/rishad004/learning-platform-go/enrollment-service/proto/client/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	Db := pkg.InitDatabase()

	connPayment, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer connPayment.Close()
	paymentService := payment_pb.NewPaymentServiceClient(connPayment)

	repo := repository.NewEnrollemntRepository(Db)
	enrollmentService := service.NewEnrollemntService(repo, paymentService)
	enrollmentHandler := grpchandlers.NewEnrollemntHandler(enrollmentService)
	httpController := httpcontrollers.NewEnrollemntController(enrollmentService)

	g := grpc.NewServer()
	pb.RegisterEnrollmentServiceServer(g, enrollmentHandler)

	reflection.Register(g)

	go func() {
		listen, err := net.Listen("tcp", ":7071")
		if err != nil {
			log.Fatal("failed to listen, err: ", err)
		}

		log.Println("Server listening on port :7071")
		if err := g.Serve(listen); err != nil {
			log.Fatal("failed to listen - err: ", err)
		}
	}()

	go func() {
		r := gin.Default()
		httpController.EnrollmentRouters(r)

		if err := r.Run(":7070"); err != nil {
			log.Fatal("failed to listen, err: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
