package enrollment

import (
	"errors"
	"log"
	"v/internal/course"
	"v/internal/domain"
	"v/internal/user"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}
	service struct {
		userSrv   user.Service
		courseSrv course.Service
		repo      Repository
	}
)

func NewService(repo Repository, userSrv user.Service, courseSrv course.Service) Service {
	return &service{
		userSrv:   userSrv,
		courseSrv: courseSrv,
		repo:      repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {
	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}
	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		return nil, errors.New("user id doesn't exists")
	}
	if _, err := s.courseSrv.Get(enroll.CourseID); err != nil {
		return nil, errors.New("course id doesn't exists")
	}
	if err := s.repo.Create(enroll); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return enroll, nil
}
