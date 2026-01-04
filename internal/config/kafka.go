package config

import (
	"context"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewKafkaConsumerGroup(config *viper.Viper, log *logrus.Logger) (*kafka.Reader, error) {
	brokers := strings.Split(config.GetString("kafka.bootstrap.servers"), ",")
	groupID := config.GetString("kafka.group.id")
	topic := config.GetString("kafka.topic")

	log.Infof("Initializing Kafka consumer with brokers: %v, groupID: %s, topic: %s", brokers, groupID, topic)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return reader, nil
}

func NewKafkaConsumer(config *viper.Viper, log *logrus.Logger, topic string) *kafka.Reader {
	brokers := strings.Split(config.GetString("kafka.bootstrap.servers"), ",")
	groupID := config.GetString("kafka.group.id")

	log.Infof("Initializing Kafka consumer with brokers: %v, groupID: %s, topic: %s", brokers, groupID, topic)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return reader
}

func ConsumeTopic(signal chan string, consumer *kafka.Reader, topic string, log *logrus.Logger, handler func(message *kafka.Message) error) {
	log.Infof("Starting to consume from topic: %s", topic)

	stop := false

	for !stop {
		select {
		case <-signal:
			log.Info("Got one of stop signals, shutting down consumer gracefully")
			stop = true
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

	log.Info("Closing consumer")
	err := consumer.Close()
	if err != nil {
		log.Errorf("Failed to close consumer: %v", err)
	}
}

func NewKafkaProducer(config *viper.Viper, log *logrus.Logger) (*kafka.Writer, error) {
	brokers := strings.Split(config.GetString("kafka.bootstrap.servers"), ",")

	log.Infof("Initializing Kafka producer with brokers: %v", brokers)

	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}

	return writer, nil
}
