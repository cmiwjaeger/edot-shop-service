package config

import (
	"edot-monorepo/services/shop-service/internal/delivery/http/controller"
	"edot-monorepo/services/shop-service/internal/delivery/http/route"
	"edot-monorepo/services/shop-service/internal/gateway/messaging"
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"
	"edot-monorepo/services/shop-service/internal/usecase"
	"edot-monorepo/shared/events"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
	Producer *kafka.Producer
}

func Bootstrap(config *BootstrapConfig) {

	shopCreatedProducer := messaging.NewShopProducer[*events.ShopCreatedEvent]("shop_created", config.Producer, config.Log)
	shopWhAssignedProducer := messaging.NewShopProducer[*events.ShopWarehouseAssignedEvent]("shop_assign_warehouse", config.Producer, config.Log)

	shopRepository := repository.NewShopRepository(config.Log)
	shopWhRepository := repository.NewShopWarehouseRepository(config.Log)

	shopBaseUseCase := usecase.NewShopUseCase(config.DB, config.Log, shopRepository, config.Validate)
	shopCreateUseCase := usecase.NewShopCreateUseCase(shopBaseUseCase, shopCreatedProducer)

	shopAssignUseCase := usecase.NewShopAssignUseCase(config.DB, config.Log, shopWhRepository, config.Validate, shopWhAssignedProducer)

	shopController := controller.NewShopController(shopCreateUseCase, shopAssignUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:            config.App,
		ShopController: shopController,
	}

	routeConfig.Setup()
}
