package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"edot-monorepo/services/shop-service/internal/config"
	"edot-monorepo/services/shop-service/internal/delivery/messaging"
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)

	logger.Info("Starting worker service")

	ctx, cancel := context.WithCancel(context.Background())

	go RunShopConsumer(logger, viperConfig, ctx)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}

func RunShopConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup shop consumer")
	db := config.NewDatabase(viperConfig, logger)
	validate := config.NewValidator(viperConfig)
	shopRepo := repository.NewShopRepository(logger)
	consumer := config.NewKafkaConsumer(viperConfig, logger)
	warehouseHandler := messaging.NewShopConsumer(db, validate, shopRepo, logger)

	messaging.ConsumeTopic(ctx, consumer, "warehouse_created", logger, warehouseHandler.ConsumeWarehouseCreated)

}
