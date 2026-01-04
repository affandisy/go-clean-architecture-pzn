package config

import (
	"strings"

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

func NewKafkaProducer(config *viper.Viper, log *logrus.Logger) (*kafka.Writer, error) {
	brokers := strings.Split(config.GetString("kafka.bootstrap.servers"), ",")

	log.Infof("Initializing Kafka producer with brokers: %v", brokers)

	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}

	return writer, nil
}
