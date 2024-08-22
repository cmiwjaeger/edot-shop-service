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
	shopAssignUseCase *usecase.ShopAssignUseCase
	Log               *logrus.Logger
	Validate          *validator.Validate
}

func NewShopController(shopCreateUseCase *usecase.ShopCreateUseCase, shopAssignUseCase *usecase.ShopAssignUseCase, log *logrus.Logger, validate *validator.Validate) *ShopController {
	return &ShopController{
		shopCreateUseCase: shopCreateUseCase,
		shopAssignUseCase: shopAssignUseCase,
		Log:               log,
		Validate:          validate,
	}
}

func (c *ShopController) Create(ctx *fiber.Ctx) error {

	request := new(model.ShopCreateRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.shopCreateUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create shop : %+v", response)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ShopCreateResponse]{Data: response})
}

func (c *ShopController) AssignWarehouse(ctx *fiber.Ctx) error {
	request := new(model.ShopAssignWarehouseRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.shopAssignUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create shop : %+v", response)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ShopAssignWarehouseRequest]{Data: request})
}
