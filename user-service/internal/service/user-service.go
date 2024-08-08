package service

import (
	"github.com/rishad004/learning-platform-go/user-service/internal/model"
	institute_pb "github.com/rishad004/learning-platform-go/user-service/proto/client/institute"
)

type UserRepo interface {
	CreateUser(user model.Users) error
	FindByUsername(user model.Users) (model.Users, error)
}

type UserService interface {
	Signup(user model.Users) error
	Login(user model.Users) (model.Users, error)
}

type userService struct {
	instituteClient institute_pb.AdminServiceClient
	repo            UserRepo
}

func NewUserService(repo UserRepo, instituteService institute_pb.AdminServiceClient) *userService {
	return &userService{repo: repo, instituteClient: instituteService}
}

func (u *userService) Signup(user model.Users) error {
	if err := u.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(user model.Users) (model.Users, error) {
	check, err := u.repo.FindByUsername(user)
	if err != nil {
		return model.Users{}, err
	}
	return check, nil
}
