package messaging

import (
	"context"
	"edot-monorepo/services/shop-service/internal/entity"
	repository "edot-monorepo/services/shop-service/internal/repository/gorm"
	"edot-monorepo/shared/events"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarehouseConsumer struct {
	DB             *gorm.DB
	Validate       *validator.Validate
	ShopRepository *repository.ShopRepository
	Log            *logrus.Logger
}

func NewShopConsumer(db *gorm.DB, validate *validator.Validate, shopRepo *repository.ShopRepository, log *logrus.Logger) *WarehouseConsumer {
	return &WarehouseConsumer{
		DB:             db,
		Validate:       validate,
		ShopRepository: shopRepo,
		Log:            log,
	}
}

func (c WarehouseConsumer) ConsumeWarehouseCreated(message *kafka.Message, ctx context.Context) error {
	event := new(events.WarehouseCreatedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		c.Log.WithError(err).Error("error unmarshalling WarehouseCreatedEvent")
		return err
	}
	data := &entity.Warehouse{
		Name:   event.Name,
		Status: event.Status,
	}

	err := c.DB.Create(data).Error
	if err != nil {
		c.Log.WithError(err).Error("error insert into db")
	}

	c.Log.Infof("Received topic warehouse with event: %v from partition %d", event, message.TopicPartition.Partition)
	return nil
}
