package repository

import (
	"edot-monorepo/services/shop-service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShopWarehouseRepository struct {
	Repository[entity.ShopWarehouse]
	Log *logrus.Logger
}

func NewShopWarehouseRepository(log *logrus.Logger) *ShopWarehouseRepository {
	return &ShopWarehouseRepository{
		Log: log,
	}
}

func (r *ShopWarehouseRepository) Manage(db *gorm.DB, data *entity.ShopWarehouse) (assigned bool, err error) {

	err = db.First(data).Error
	if err == nil {
		err = db.Delete(data).Error
		assigned = false
	} else {
		err = db.Create(data).Error
		assigned = true
	}

	return
}
