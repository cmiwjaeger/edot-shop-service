package entity

import (
	"time"

	"github.com/google/uuid"
)

// Shop is a struct that represents a shop entity
type Shop struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
}

func (u *Shop) TableName() string {
	return "shops"
}
