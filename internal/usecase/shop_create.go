package usecase

import (
	"context"
	"edot-monorepo/services/shop-service/internal/entity"
	"edot-monorepo/services/shop-service/internal/model"
	"edot-monorepo/services/shop-service/internal/model/converter"

	"github.com/gofiber/fiber/v2"
)

type ShopCreateUseCase struct {
	*ShopBaseUseCase
}

func NewShopCreateUseCase(shopBaseUseCase *ShopBaseUseCase) *ShopCreateUseCase {
	return &ShopCreateUseCase{
		shopBaseUseCase,
	}
}

func (c *ShopCreateUseCase) Exec(ctx context.Context, request *model.ShopCreateRequest) (*model.ShopCreateResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	shop := &entity.Shop{
		Name:    request.Name,
		Address: request.Address,
	}
	if err := c.ShopRepository.Create(tx, shop); err != nil {
		c.Log.Warnf("Failed create shop to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	event := converter.ShopToEvent(shop)
	if err = c.Producer.Produce(ctx, "shop_created", event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ShopToResponse(shop), nil
}
