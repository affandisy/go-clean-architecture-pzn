package messaging

import (
	"go-clean-architecture-pzn/internal/model"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type AddressProducer struct {
	Producer[*model.AddressEvent]
}

func NewAddressProducer(producer *kafka.Writer, log *logrus.Logger) *AddressProducer {
	return &AddressProducer{
		Producer: Producer[*model.AddressEvent]{
			Producer: producer,
			Topic:    "addresses",
			Log:      log,
		},
	}
}
