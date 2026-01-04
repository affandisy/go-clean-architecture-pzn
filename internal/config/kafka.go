package config

import (
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func NewKafkaConsumerGroup(config *viper.Viper) (*kafka.Reader, error) {
	brokers := strings.Split(config.GetString("kafka.bootstrap.servers"), ",")
	groupID := config.GetString("kafka.group.id")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   config.GetString("kafka.topic"),
	})

	return reader, nil
}

func NewKafkaProducer(config *viper.Viper) (*kafka.Writer, error) {
	brokers := strings.Split(config.GetString("kafka.bootstrap.servers"), ",")

	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}

	return writer, nil
}
