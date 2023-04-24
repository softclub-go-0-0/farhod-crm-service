package models

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" binding:"required"`
	Surname   string         `json:"surname" binding:"required"`
	Phone     string         `json:"phone" binding:"required,numeric"`
	Email     string         `json:"email" binding:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Groups []Group `json:"groups,omitempty"`
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
