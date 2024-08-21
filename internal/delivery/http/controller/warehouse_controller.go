package controller

import (
	"edot-monorepo/services/shop-service/internal/model"
	"edot-monorepo/services/shop-service/internal/usecase"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ShopController struct {
	shopCreateUseCase *usecase.ShopCreateUseCase
	Log               *logrus.Logger
	Validate          *validator.Validate
}

func NewShopController(shopCreateUseCase *usecase.ShopCreateUseCase, log *logrus.Logger, validate *validator.Validate) *ShopController {
	return &ShopController{
		shopCreateUseCase: shopCreateUseCase,
		Log:               log,
		Validate:          validate,
	}
}

func (c *ShopController) Create(ctx *fiber.Ctx) error {

	return ctx.JSON(model.WebResponse[*model.ShopResponse]{Data: nil})
}
