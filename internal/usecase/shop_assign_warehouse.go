package usecase

import (
	"context"
	"edot-monorepo/services/shop-service/internal/entity"
	"edot-monorepo/services/shop-service/internal/gateway/messaging"
	"edot-monorepo/services/shop-service/internal/model"
	"edot-monorepo/services/shop-service/internal/model/converter"
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShopAssignUseCase struct {
	DB           *gorm.DB
	Log          *logrus.Logger
	ShopWhRepo   *repository.ShopWarehouseRepository
	Validate     *validator.Validate
	ShopProducer *messaging.ShopProducer[model.Event]
}

func NewShopAssignUseCase(DB *gorm.DB, Log *logrus.Logger, ShopWhRepo *repository.ShopWarehouseRepository, Validate *validator.Validate, ShopProducer *messaging.ShopProducer[model.Event]) *ShopAssignUseCase {
	return &ShopAssignUseCase{
		DB:           DB,
		Log:          Log,
		ShopWhRepo:   ShopWhRepo,
		Validate:     Validate,
		ShopProducer: ShopProducer,
	}
}

func (c *ShopAssignUseCase) Exec(ctx context.Context, request *model.ShopAssignWarehouseRequest) (*model.ShopAssignWarehouseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	shopWh := &entity.ShopWarehouse{
		ShopID:      request.ShopID,
		WarehouseID: request.WarehouseID,
	}

	assigned, err := c.ShopWhRepo.Manage(tx, shopWh)
	if err != nil {
		c.Log.Warnf("Failed assign shop to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	event := converter.ShopWhToEvent(shopWh, assigned)
	if err := c.ShopProducer.SendAsync(event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ShopWhToResponse(shopWh), nil

}
