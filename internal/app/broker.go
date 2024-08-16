package app

import (
	"awesomeProject/internal/config"
	"fmt"
	"github.com/IBM/sarama"
)

func InitBroker(cfg *config.Broker) (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("localhost:%d", cfg.Port)}, nil)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
