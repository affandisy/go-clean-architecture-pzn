package main

import (
	"go-clean-architecture-pzn/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// fmt.Println("Worker")
	// viperConfig, err := config.NewViper()
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %w", err))
	// }

	// consumerUser, err := config.NewKafkaConsumerGroup(viperConfig, logger)
	// signalUser := make(chan string)

	// consumer, err := config.NewKafkaConsumerGroup(viperConfig, logger)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error kafka consumer: %w", err))
	// }

	// err = consumer.Close()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error close kafka consumer: %w", err))
	// }

	consumerUser := config.NewKafkaConsumer(viperConfig, logger, "users")
	signalUser := make(chan string)

	consumerContact := config.NewKafkaConsumer(viperConfig, logger, "contacts")
	signalContact := make(chan string)

	consumerAddress := config.NewKafkaConsumer(viperConfig, logger, "addresses")
	signalAddress := make(chan string)

	go config.ConsumeTopic(signalUser, consumerUser, "users", logger, func(message *kafka.Message) error {
		event := string(message.Value)
		logger.Infof("Received topic users with event: %s from partition %d", event, message.Partition)
		return nil
	})

	go config.ConsumeTopic(signalContact, consumerContact, "contacts", logger, func(message *kafka.Message) error {
		event := string(message.Value)
		logger.Infof("Received topic contacts with event: %s from partition %d", event, message.Partition)
		return nil
	})

	go config.ConsumeTopic(signalAddress, consumerAddress, "addresses", logger, func(message *kafka.Message) error {
		event := string(message.Value)
		logger.Infof("Received topic addresses with event: %s from partition %d", event, message.Partition)
		return nil
	})

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	s := <-terminateSignals
	logger.Info("Got one of stop signals, shutting down server gracefully, SIGNAL NAME :", s)
	signalUser <- "stop"
	signalContact <- "stop"
	signalAddress <- "stop"

	time.Sleep(5 * time.Second)
}
