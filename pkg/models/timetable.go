package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Timetable struct {
	ID        uuid.UUID      `json:"id" gorm:"primarykey"`
	Classroom string         `json:"classroom" binding:"required"`
	Start     Time           `json:"start" binding:"required"`
	Finish    Time           `json:"finish" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (t *Timetable) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID, err = uuid.NewUUID()
	return err
}
