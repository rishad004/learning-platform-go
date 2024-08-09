package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	grpchandlers "github.com/rishad004/learning-platform-go/payment-service/internal/handler/grpc-handlers"
	httpcontrollers "github.com/rishad004/learning-platform-go/payment-service/internal/handler/http-controllers"
	"github.com/rishad004/learning-platform-go/payment-service/internal/repository"
	"github.com/rishad004/learning-platform-go/payment-service/internal/service"
	"github.com/rishad004/learning-platform-go/payment-service/pkg"
	pb "github.com/rishad004/learning-platform-go/payment-service/proto"
	institute_pb "github.com/rishad004/learning-platform-go/payment-service/proto/client/institute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	Db := pkg.InitDatabase()

	connInstitute, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer connInstitute.Close()
	instituteService := institute_pb.NewAdminServiceClient(connInstitute)

	repo := repository.NewPaymentRepository(Db)
	paymentService := service.NewPaymentService(repo, instituteService)
	paymentHandler := grpchandlers.NewPaymentHandler(paymentService)
	httpController := httpcontrollers.NewPaymentController(paymentService)

	g := grpc.NewServer()
	pb.RegisterPaymentServiceServer(g, paymentHandler)

	reflection.Register(g)

	go func() {
		listen, err := net.Listen("tcp", ":9091")
		if err != nil {
			log.Fatal("failed to listen, err: ", err)
		}

		log.Println("Server listening on port :9091")
		if err := g.Serve(listen); err != nil {
			log.Fatal("failed to listen - err: ", err)
		}
	}()

	go func() {
		r := gin.Default()
		httpController.PaymentRouters(r)
		r.LoadHTMLGlob("../template/*")

		if err := r.Run(":9090"); err != nil {
			log.Fatal("failed to listen, err: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
