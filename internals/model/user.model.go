package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email    string    `gorm:"unique;not null"`
	Username string    `gorm:"unique;not null"`
	FullName string    `gorm:"not null"`
	UserRole Role      `gorm:"default:user"`
}
