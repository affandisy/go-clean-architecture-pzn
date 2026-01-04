package messaging

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func ConsumeTopic(ctx context.Context, consumer *kafka.Reader, topic string, log *logrus.Logger, handler func(message *kafka.Message) error) {
	log.Infof("Starting to consume from topic: %s", topic)

	// stop := false
	run := true

	for run {
		select {
		case <-ctx.Done():
			// log.Info("Got one of stop signals, shutting down consumer gracefully")
			run = false
		default:
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			message, err := consumer.ReadMessage(ctx)
			cancel()
			if err == nil {
				log.Debugf("Received message from topic %s: partition=%d offset=%d", topic, message.Partition, message.Offset)

				err := handler(&message)
				if err != nil {
					log.Errorf("Failed to process message from partition %d offset %d: %v", message.Partition, message.Offset, err)
				} else {
					log.Debugf("Successfully processed message from partition %d offset %d", message.Partition, message.Offset)
				}
			} else {
				// Check if error is timeout (which is normal)
				if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Timeout() {
					// Timeout is expected, continue
					continue
				}
				log.Warnf("Consumer error while reading message: %v", err)
			}
		}
	}

	log.Infof("Closing consumer for topic: %s", topic)
	err := consumer.Close()
	if err != nil {
		log.Errorf("Failed to close consumer: %v", err)
	}
}
