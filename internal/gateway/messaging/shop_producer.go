package messaging

import (
	"edot-monorepo/services/shop-service/internal/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type ShopProducer[T model.Event] struct {
	Producer[T]
}

func NewShopProducer[T model.Event](topic string, producer *kafka.Producer, log *logrus.Logger) *ShopProducer[model.Event] {

	return &ShopProducer[model.Event]{
		Producer: Producer[model.Event]{
			Producer: producer,
			Topic:    topic,
			Log:      log,
		},
	}
}
