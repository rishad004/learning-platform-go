package service

type UserRepo interface {
}

type UserService interface {
}

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *userService {
	return &userService{repo: repo}
}
