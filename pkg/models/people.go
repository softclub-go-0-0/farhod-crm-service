package models

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID uint `gorm:"primarykey"`

	Name    string
	Surname string
	Phone   string
	Email   string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Groups []Group
}

type Student struct {
	ID uint `gorm:"primarykey"`

	GroupID uint
	Name    string
	Surname string
	Phone   string
	Email   string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Group Group
}
