package usecase

import (
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShopBaseUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	ShopRepository *repository.ShopRepository
	Validate       *validator.Validate
}

func NewShopUseCase(db *gorm.DB, log *logrus.Logger, shopRepo *repository.ShopRepository, validate *validator.Validate) *ShopBaseUseCase {
	return &ShopBaseUseCase{
		DB:             db,
		Log:            log,
		ShopRepository: shopRepo,
		Validate:       validate,
	}
}
