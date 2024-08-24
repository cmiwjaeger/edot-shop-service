package usecase

import (
	"edot-monorepo/services/shop-service/internal/gateway/messaging"
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
	Producer       *messaging.Producer
}

func NewShopUseCase(db *gorm.DB, log *logrus.Logger, shopRepo *repository.ShopRepository, validate *validator.Validate, producer *messaging.Producer) *ShopBaseUseCase {
	return &ShopBaseUseCase{
		DB:             db,
		Log:            log,
		ShopRepository: shopRepo,
		Validate:       validate,
		Producer:       producer,
	}
}
