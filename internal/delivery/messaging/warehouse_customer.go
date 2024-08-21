package messaging

import (
	"edot-monorepo/services/shop-service/internal/model"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type ShopConsumer struct {
	Log *logrus.Logger
}

func NewShopConsumer(log *logrus.Logger) *ShopConsumer {
	return &ShopConsumer{
		Log: log,
	}
}

func (c ShopConsumer) Consume(message *kafka.Message) error {
	ContactEvent := new(model.Shop)
	if err := json.Unmarshal(message.Value, ContactEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Contact event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic contacts with event: %v from partition %d", ContactEvent, message.TopicPartition.Partition)
	return nil
}
