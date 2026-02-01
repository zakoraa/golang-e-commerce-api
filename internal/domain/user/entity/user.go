package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`

	Role UserRole `gorm:"type:user_role;not null"`

	CreatedAt time.Time
	CreatedBy *uuid.UUID

	UpdatedAt *time.Time
	UpdatedBy *uuid.UUID

	DeletedAt gorm.DeletedAt
	DeletedBy *uuid.UUID
}

func (u *User) BeforeCreate(tx any) error { u.ID = uuid.New() return nil }
