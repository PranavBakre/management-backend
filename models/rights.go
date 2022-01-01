package models

import "github.com/google/uuid"

type Right struct {
	ID    uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Right string    `gorm:"not null;unique;"`
}

func (Right) TableName() string {
	return "rights"
}
