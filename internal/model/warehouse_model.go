package model

import (
	"time"

	"github.com/google/uuid"
)

type Shop struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShopResponse struct {
	ID uuid.UUID `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ShopCreateRequest struct {
	ID uuid.UUID
}
