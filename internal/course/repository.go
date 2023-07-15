package course

import (
	"fmt"
	"log"
	"strings"
	"time"
	"v/internal/domain"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		GetAll(filters Filters, offsite, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		Update(id string, name, startDate, endDate *string) error
		Count(filters Filters) (int, error)
		Delete(id string) error
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

func (r *repo) Create(course *domain.Course) error {
	if err := r.db.Create(course).Error; err != nil {
		log.Fatal("error:", err)
	}

	return nil
}

func (r *repo) GetAll(filters Filters, offsite, limit int) ([]domain.Course, error) {
	var c []domain.Course
	tx := r.db.Model(&c)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offsite)
	result := tx.Order("created_at desc").Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

func (r *repo) Get(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}
	if tx := r.db.First(&course).Error; tx != nil {
		return nil, tx
	}
	return &course, nil
}

func (r *repo) Update(id string, name, startDate, endDate *string) error {
	newStartDate, err := time.Parse("2006-01-02", *startDate)
	if err != nil {
		return err
	}
	newEndDate, err := time.Parse("2006-01-02", *endDate)
	if err != nil {
		return err
	}
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = name
	}
	if startDate != nil {
		values["start_date"] = newStartDate
	}
	if endDate != nil {
		values["end_date"] = newEndDate
	}
	if err := r.db.Model(&domain.Course{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
func (repo *repo) Delete(id string) error {
	course := domain.Course{ID: id}
	if err := repo.db.Delete(&course).Error; err != nil {
		return err
	}
	return nil
}
func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(first_name) like ?", filters.Name)
	}
	return tx
}
