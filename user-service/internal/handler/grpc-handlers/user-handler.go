package grpchandlers

import (
	"github.com/rishad004/learning-platform-go/user-service/internal/service"
	pb "github.com/rishad004/learning-platform-go/user-service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	Service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}
