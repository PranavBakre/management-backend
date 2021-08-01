package models

import (
	"time"

	"github.com/m4rw3r/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uuid.UUID      `gorm:"default:uuid_generate_v4()" json:"id" form:"id"`
	Name      string         `gorm:"default:null;not null" json:"name" form:"name"`
	Email     string         `gorm:"default:null;not null" json:"email" form:"email"`
	GoogleID  *string        `gorm:"index;unique" json:"google_id" form:"google_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName for User
func (User) TableName() string {
	return "users"
}
