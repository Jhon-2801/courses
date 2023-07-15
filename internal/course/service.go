package course

import (
	"log"
	"time"
	"v/internal/domain"
)

type (
	Filters struct {
		Name string
	}
	Service interface {
		Create(name, startDate, endDate string) (*domain.Course, error)
		GetAll(filters Filters, offsite, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		UpDate(id string, name, startDate, endDate *string) error
		Count(filters Filters) (int, error)
		Delete(id string) error
	}
	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(name, startDate, endDate string) (*domain.Course, error) {
	newStartDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	newEndDate, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	course := &domain.Course{
		Name:      name,
		StartDate: newStartDate,
		EndDate:   newEndDate,
	}
	err = s.repo.Create(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s service) GetAll(filters Filters, offsite, limit int) ([]domain.Course, error) {
	courses, err := s.repo.GetAll(filters, offsite, limit)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (s service) Get(id string) (*domain.Course, error) {
	course, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s service) UpDate(id string, name, startDate, endDate *string) error {
	err := s.repo.Update(id, name, startDate, endDate)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
