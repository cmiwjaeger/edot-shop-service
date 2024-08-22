package entity

import (
	"github.com/google/uuid"
)

// Shop is a struct that represents a shop entity
type ShopWarehouse struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ShopID      uuid.UUID `gorm:"column:shop_id"`
	WarehouseID uuid.UUID `gorm:"column:warehouse_id"`
}

func (u *ShopWarehouse) TableName() string {
	return "shop_warehouses"
}
