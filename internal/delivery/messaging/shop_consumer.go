package messaging

import (
	"context"
	"edot-monorepo/services/shop-service/internal/entity"
	"edot-monorepo/shared/events"

	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarehouseConsumer struct {
	Log       *logrus.Logger
	DB        *gorm.DB
	Validator *validator.Validate
}

func NewShopConsumer(log *logrus.Logger, db *gorm.DB, validate *validator.Validate) *WarehouseConsumer {
	return &WarehouseConsumer{
		Log:       log,
		DB:        db,
		Validator: validate,
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

	c.Log.Infof("Received topic warehouse with event: %v from partition %s", event, message.Topic)
	return nil
}
