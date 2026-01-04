package messaging

import (
	"encoding/json"
	"go-clean-architecture-pzn/internal/model"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type AddressConsumer struct {
	Log *logrus.Logger
}

func NewAddressConsumer(log *logrus.Logger) *AddressConsumer {
	return &AddressConsumer{
		Log: log,
	}
}

func (c AddressConsumer) Consume(message *kafka.Message) error {
	AddressEvent := new(model.AddressEvent)
	if err := json.Unmarshal(message.Value, AddressEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Address event")
		return err
	}

	c.Log.Infof("Received topic Address with event: %v from partition %d", AddressEvent, message.Partition)
	return nil
}
