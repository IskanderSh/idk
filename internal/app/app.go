package app

import (
	"awesomeProject/internal/config"
	"github.com/IBM/sarama"
	"log/slog"
)

type Service struct {
	MessageProducer sarama.SyncProducer
}

func StartService(log *slog.Logger, cfg *config.Config) error {
	var service Service

	// "message" topic producer
	messageProducer, err := InitBroker(&cfg.Broker)
	if err != nil {
		return err
	}
	service.MessageProducer = messageProducer

	return nil
}
