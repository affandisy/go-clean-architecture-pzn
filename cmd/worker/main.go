package main

import (
	"context"
	"go-clean-architecture-pzn/internal/config"
	"go-clean-architecture-pzn/internal/delivery/messaging"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// fmt.Println("Worker")
	// viperConfig, err := config.NewViper()
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting Worker Service")

	ctx, cancel := context.WithCancel(context.Background())
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

	// logger.Info("setup user consumer")
	// consumerUser := config.NewKafkaConsumer(viperConfig, logger, "users")
	// signalUser := make(chan string)
	// logger.Info("setup user consumer")
	// userConsumer := config.NewKafkaConsumer(viperConfig, logger, "users")
	// // userSignal := make(chan string)
	// userHandler := messaging.NewUserConsumer(logger)
	// go messaging.ConsumeTopic(ctx, userConsumer, "users", logger, userHandler.Consume)

	// // consumerContact := config.NewKafkaConsumer(viperConfig, logger, "contacts")
	// // signalContact := make(chan string)

	// logger.Info("setup contact consumer")
	// contactConsumer := config.NewKafkaConsumer(viperConfig, logger, "contacts")
	// // contactSignal := make(chan string)
	// contactHandler := messaging.NewContactConsumer(logger)
	// go messaging.ConsumeTopic(ctx, contactConsumer, "contacts", logger, contactHandler.Consume)

	// // consumerAddress := config.NewKafkaConsumer(viperConfig, logger, "addresses")
	// // signalAddress := make(chan string)

	// logger.Info("setup address consumer")
	// addressConsumer := config.NewKafkaConsumer(viperConfig, logger, "addresses")
	// // addressSignal := make(chan string)
	// addressHandler := messaging.NewAddressConsumer(logger)
	// go messaging.ConsumeTopic(ctx, addressConsumer, "addresses", logger, addressHandler.Consume)

	// go config.ConsumeTopic(signalUser, consumerUser, "users", logger, func(message *kafka.Message) error {
	// 	event := string(message.Value)
	// 	logger.Infof("Received topic users with event: %s from partition %d", event, message.Partition)
	// 	return nil
	// })

	// go config.ConsumeTopic(signalContact, consumerContact, "contacts", logger, func(message *kafka.Message) error {
	// 	event := string(message.Value)
	// 	logger.Infof("Received topic contacts with event: %s from partition %d", event, message.Partition)
	// 	return nil
	// })

	// go config.ConsumeTopic(signalAddress, consumerAddress, "addresses", logger, func(message *kafka.Message) error {
	// 	event := string(message.Value)
	// 	logger.Infof("Received topic addresses with event: %s from partition %d", event, message.Partition)
	// 	return nil
	// })

	go RunUserConsumer(logger, viperConfig, ctx)
	go RunContactConsumer(logger, viperConfig, ctx)
	go RunAddressConsumer(logger, viperConfig, ctx)

	logger.Info("Worker running")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	s := <-terminateSignals
	logger.Info("Got one of stop signals, shutting down consumer gracefully, SIGNAL NAME :", s)
	// userSignal <- "stop"
	// contactSignal <- "stop"
	// addressSignal <- "stop"
	cancel()

	// s := <-terminateSignals
	// logger.Info("Got one of stop signals, shutting down server gracefully, SIGNAL NAME :", s)
	// signalUser <- "stop"
	// signalContact <- "stop"
	// signalAddress <- "stop"

	time.Sleep(5 * time.Second)
}

func RunAddressConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup address consumer")
	// addressConsumer := config.NewKafkaConsumer(viperConfig, logger, "addresses")
	addressConsumerGroup := config.NewKafkaConsumerGroup(viperConfig, logger)
	addressHandler := messaging.NewAddressConsumer(logger)
	messaging.ConsumeTopic(ctx, addressConsumerGroup, "addresses", logger, addressHandler.Consume)
}

func RunContactConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup contact consumer")
	// contactConsumer := config.NewKafkaConsumer(viperConfig, logger, "contacts")
	contactConsumerGroup := config.NewKafkaConsumerGroup(viperConfig, logger)
	contactHandler := messaging.NewContactConsumer(logger)
	messaging.ConsumeTopic(ctx, contactConsumerGroup, "contacts", logger, contactHandler.Consume)
}

func RunUserConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup user consumer")
	// userConsumer := config.NewKafkaConsumer(viperConfig, logger, "users")
	userConsumerGroup := config.NewKafkaConsumerGroup(viperConfig, logger)
	userHandler := messaging.NewUserConsumer(logger)
	messaging.ConsumeTopic(ctx, userConsumerGroup, "users", logger, userHandler.Consume)
}
