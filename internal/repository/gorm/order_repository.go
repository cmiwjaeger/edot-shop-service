package repository

import (
	"edot-monorepo/services/shop-service/internal/entity"

	"github.com/sirupsen/logrus"
)

type ShopRepository struct {
	Repository[entity.Shop]
	Log *logrus.Logger
}

func NewShopRepository(log *logrus.Logger) *ShopRepository {
	return &ShopRepository{
		Log: log,
	}
}
