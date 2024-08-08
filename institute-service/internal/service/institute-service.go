package service

import "github.com/rishad004/learning-platform-go/institute-service/internal/model"

type InstituteRepo interface {
	FindByAdminname(admin model.Amdin) (model.Amdin, error)
	CreateCourse(course model.Courses) error
	FetchCourses() ([]model.Courses, error)
}

type InstituteService interface {
	Login(admin model.Amdin) (model.Amdin, error)
	AddCourse(course model.Courses) error
	GetCourseInfo() ([]model.Courses, error)
}

type instituteService struct {
	repo InstituteRepo
}

func NewInstituteService(repo InstituteRepo) InstituteService {
	return &instituteService{repo: repo}
}

func (s *instituteService) Login(admin model.Amdin) (model.Amdin, error) {
	check, err := s.repo.FindByAdminname(admin)
	if err != nil {
		return model.Amdin{}, err
	}
	return check, nil
}

func (s *instituteService) AddCourse(course model.Courses) error {
	if err := s.repo.CreateCourse(course); err != nil {
		return err
	}
	return nil
}

func (s *instituteService) GetCourseInfo() ([]model.Courses, error) {
	courses, err := s.repo.FetchCourses()
	if err != nil {
		return nil, err
	}
	return courses, nil
}
