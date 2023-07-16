package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id,omitempty"`
	User      *User      `json:"user,omitempty"`
	CourseID  string     `json:"course_id"`
	Course    *Course    `json:"course,omitempty"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (c *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
