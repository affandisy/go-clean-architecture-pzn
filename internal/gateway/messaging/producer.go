package messaging

import (
	"context"
	"encoding/json"
	"go-clean-architecture-pzn/internal/model"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Producer[T model.Event] struct {
	// Producer *kafka.Writer
	Producer sarama.SyncProducer
	Topic    string
	Log      *logrus.Logger
}

func (p *Producer[T]) Send(ctx context.Context, event T) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: p.Topic,
		Key:   sarama.StringEncoder(event.GetId()),
		Value: sarama.ByteEncoder(jsonData),
	}

	partition, offset, err := p.Producer.SendMessage(message)
	if err != nil {
		p.Log.WithError(err).Error("failed to produce message")
		return err
	}

	p.Log.Debugf("Message sent to topic %s, partition %d, offset %d", p.Topic, partition, offset)
	return nil
}
