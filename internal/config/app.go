package config

import (
	"edot-monorepo/services/shop-service/internal/delivery/http/controller"
	"edot-monorepo/services/shop-service/internal/delivery/http/route"
	"edot-monorepo/services/shop-service/internal/gateway/messaging"
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"
	"edot-monorepo/services/shop-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
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

	Reader *kafka.Reader
	Writer *kafka.Writer
}

func Bootstrap(config *BootstrapConfig) {

	producer := messaging.NewProducer(config.Writer, config.Log)

	shopRepository := repository.NewShopRepository(config.Log)
	shopWhRepository := repository.NewShopWarehouseRepository(config.Log)

	shopBaseUseCase := usecase.NewShopUseCase(config.DB, config.Log, shopRepository, config.Validate, producer)
	shopCreateUseCase := usecase.NewShopCreateUseCase(shopBaseUseCase)

	shopAssignUseCase := usecase.NewShopAssignUseCase(config.DB, config.Log, shopWhRepository, config.Validate, producer)

	shopController := controller.NewShopController(shopCreateUseCase, shopAssignUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:            config.App,
		ShopController: shopController,
	}

	routeConfig.Setup()
}
