package usecase

import (
	"context"
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

func (c *ShopCreateUseCase) Exec(ctx context.Context, request *model.ShopCreateRequest) (*model.ShopResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.ShopRepository.Create(tx, nil); err != nil {
		c.Log.Warnf("Failed create shop to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ShopToResponse(nil), nil
}
