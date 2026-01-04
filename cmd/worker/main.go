package main

import (
	"fmt"
	"go-clean-architecture-pzn/internal/config"
)

func main() {
	// fmt.Println("Worker")
	// viperConfig, err := config.NewViper()
	viperConfig := config.NewViper()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %w", err))
	// }

	consumer, err := config.NewKafkaConsumerGroup(viperConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error kafka consumer: %w", err))
	}

	err = consumer.Close()
	if err != nil {
		panic(fmt.Errorf("Fatal error close kafka consumer: %w", err))
	}
}
