package service

import "github.com/rishad004/learning-platform-go/institute-service/internal/model"

type InstituteRepo interface {
	FindByAdminname(admin model.Amdin) error
}

type InstituteService interface {
	Login(admin model.Amdin) error
}

type instituteService struct {
	repo InstituteRepo
}

func NewInstituteService(repo InstituteRepo) InstituteService {
	return &instituteService{repo: repo}
}

func (s *instituteService) Login(admin model.Amdin) error {
	if err := s.repo.FindByAdminname(admin); err != nil {
		return err
	}
	return nil
}
