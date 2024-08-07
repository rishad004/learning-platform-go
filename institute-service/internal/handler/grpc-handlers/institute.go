package grpchandlers

import (
	"github.com/rishad004/learning-platform-go/institute-service/internal/service"
	pb "github.com/rishad004/learning-platform-go/institute-service/proto"
)

type InstituteHandler struct {
	pb.UnimplementedAdminServiceServer
	Service service.InstituteService
}

func NewInstituteHandler(svc service.InstituteService) *InstituteHandler {
	return &InstituteHandler{Service: svc}
}
