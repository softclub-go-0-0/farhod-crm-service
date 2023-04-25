package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	Title      string         `json:"title" binding:"required"`
	MonthlyFee uint           `json:"monthly_fee" binding:"omitempty,number"`
	Duration   uint           `json:"duration" binding:"omitempty,number"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	Groups []Group `json:"groups"`
}

type Timetable struct {
	ID        uint           `json:"id" gorm:"primarykey"`         // 22
	Classroom string         `json:"classroom" binding:"required"` // "96"
	Start     Time           `json:"start" binding:"required"`
	Finish    Time           `json:"finish" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Group struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CourseID    uint           `json:"course_id"`
	TeacherID   uint           `json:"teacher_id"`
	TimetableID uint           `json:"timetable_id"`
	Title       string         `json:"title"`
	StartDate   time.Time      `json:"start_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	Course    *Course    `json:"course,omitempty"`
	Teacher   *Teacher   `json:"teacher,omitempty"`
	Timetable *Timetable `json:"timetable,omitempty"`
	Students  []Student  `json:"students,omitempty"`
}
