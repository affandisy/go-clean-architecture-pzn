package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"go-clean-architecture-pzn/internal/model"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer[T model.Event] struct {
	Producer *kafka.Writer
	Topic    string
	Log      *logrus.Logger
}

func (p *Producer[T]) Send(ctx context.Context, event T) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	message := kafka.Message{
		Topic: p.Topic,
		Key:   []byte(event.GetId()),
		Value: jsonData,
	}

	err = p.Producer.WriteMessages(ctx, message)
	if err != nil {
		p.Log.Errorf("Failed to send message to topic %s: %v", p.Topic, err)
		return fmt.Errorf("failed to send message: %w", err)
	}

	p.Log.Infof("Successfully sent message to topic %s with key %s", p.Topic, event.GetId())
	return nil
}
