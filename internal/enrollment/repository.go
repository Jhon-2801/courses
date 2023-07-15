package enrollment

import (
	"log"
	"v/internal/domain"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}
	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(enroll *domain.Enrollment) error {
	if err := r.db.Create(enroll).Error; err != nil {
		log.Fatal("error:", err)
		return err
	}
	return nil
}
