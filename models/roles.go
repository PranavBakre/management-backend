package models

import "github.com/google/uuid"

type Role struct {
	ID     uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Role   string    `gorm:"not null;index;unique;"`
	Rights []Right   `gorm:"default:null;many2many:role_right;"`
}

func (Role) TableName() string {
	return "roles"
}
