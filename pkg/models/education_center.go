package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID uint `gorm:"primarykey"`

	Title      string
	MonthlyFee uint
	Duration   uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Groups []Group
}

type TimeTable struct {
	ID uint `gorm:"primarykey"`

	Classroom string
	Start     time.Time
	Finish    time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Group struct {
	ID uint `gorm:"primarykey"`

	CourseID    uint
	TeacherID   uint
	TimetableID uint
	Title       string
	StartDate   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Course    Course
	Teacher   Teacher
	TimeTable TimeTable
	Students  []Student
}
