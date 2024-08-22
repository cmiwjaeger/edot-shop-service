package model

import (
	"time"

	"github.com/google/uuid"
)

type ShopCreateResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string
	Address   string
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ShopCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ShopAssignWarehouseRequest struct {
	ShopID      uuid.UUID `json:"shop_id"`
	WarehouseID uuid.UUID `json:"warehouse_id"`
}

type ShopAssignWarehouseResponse struct {
	ShopID      uuid.UUID `json:"shop_id"`
	WarehouseID uuid.UUID `json:"warehouse_id"`
}
