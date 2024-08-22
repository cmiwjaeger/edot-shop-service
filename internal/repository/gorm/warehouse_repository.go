package repository

import (
	"edot-monorepo/services/shop-service/internal/entity"

	"github.com/sirupsen/logrus"
)

type WarehouseRepository struct {
	Repository[entity.Warehouse]
	Log *logrus.Logger
}

func NewWarehouseRepository(log *logrus.Logger) *WarehouseRepository {
	return &WarehouseRepository{
		Log: log,
	}
}
