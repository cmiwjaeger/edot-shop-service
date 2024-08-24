package main

import (
	"context"
	"edot-monorepo/services/shop-service/internal/config"
	"edot-monorepo/services/shop-service/internal/delivery/messaging"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	// Initialize configurations, logger, and other dependencies
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logger)
	validate := config.NewValidator(viperConfig)

	kafkaReader := config.NewKafkaReader(viperConfig, logger)

	// Start the service
	logger.Info("Starting worker service")

	// Set up context with cancel for graceful shutdown using signal.NotifyContext
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop() // Ensure stop is called on exit

	// Use a WaitGroup to wait for the consumer to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Run the product consumer in a separate goroutine
	go func() {
		defer wg.Done()
		runConsumer(ctx, kafkaReader, logger, db, validate)
	}()

	// Wait for context cancellation (signal received)
	<-ctx.Done()
	logger.Info("Received shutdown signal, waiting for goroutines to finish")

	// Wait for the consumer to finish processing
	wg.Wait()
	logger.Info("Worker service has shut down gracefully")
}

func runConsumer(ctx context.Context, reader *kafka.Reader, logger *logrus.Logger, db *gorm.DB, validate *validator.Validate) {
	logger.Info("setup warehouse consumer")
	consumer := messaging.NewConsumer(reader)
	handler := messaging.NewShopConsumer(logger, db, validate)

	topicHandlers := map[string]messaging.ConsumerHandler{
		"warehouse_created": func(ctx context.Context, msg *kafka.Message) error {
			return handler.ConsumeWarehouseCreated(msg, ctx)
		},
	}

	consumer.Consume(ctx, topicHandlers, logger)

}
