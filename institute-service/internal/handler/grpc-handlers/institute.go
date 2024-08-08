package grpchandlers

import (
	"context"

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

func (h *InstituteHandler) GetCourseInfo(c context.Context, req *pb.GetCourseInfoRequest) (*pb.CourseInfoResponse, error) {
	var res []*pb.Course

	courses, err := h.Service.GetCourseInfo()
	if err != nil {
		return nil, err
	}
	for _, v := range courses {
		res = append(res, &pb.Course{
			Id:     int32(v.ID),
			Course: v.Name,
			Price:  int32(v.Price),
		})
	}
	return &pb.CourseInfoResponse{Courses: res}, nil
}
