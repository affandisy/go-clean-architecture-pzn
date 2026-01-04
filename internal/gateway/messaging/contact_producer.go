package messaging

import (
	"go-clean-architecture-pzn/internal/model"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type ContactProducer struct {
	Producer[*model.ContactEvent]
}

func NewContactProducer(producer *kafka.Writer, log *logrus.Logger) *ContactProducer {
	return &ContactProducer{
		Producer: Producer[*model.ContactEvent]{
			Producer: producer,
			Topic:    "contacts",
			Log:      log,
		},
	}
}
