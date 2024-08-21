package config

import (
	"edot-monorepo/services/shop-service/internal/delivery/http/controller"
	"edot-monorepo/services/shop-service/internal/delivery/http/route"
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"
	"edot-monorepo/services/shop-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	shopRepository := repository.NewShopRepository(config.Log)
	shopBaseUseCase := usecase.NewShopUseCase(config.DB, config.Log, shopRepository, config.Validate)
	shopCreateUseCase := usecase.NewShopCreateUseCase(shopBaseUseCase)

	shopController := controller.NewShopController(shopCreateUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:            config.App,
		ShopController: shopController,
	}

	routeConfig.Setup()
}
