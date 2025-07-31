package kafka

import (
	"context"
	"fmt"
	"hexabank/services/notification/domain/port"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type NotificationKafka struct {
	topic             string
	brokers           []string
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer

	notificationService port.NotificationService
}

func NewNotificationKafka(brokers []string, topic string, notificationService port.NotificationService) (*NotificationKafka, error) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return nil, fmt.Errorf("failed to start partition consumer: %w", err)
	}

	return &NotificationKafka{
		topic:             topic,
		brokers:           brokers,
		consumer:          consumer,
		partitionConsumer: partitionConsumer,

		notificationService: notificationService,
	}, nil
}

func (nk *NotificationKafka) StartConsumer() {
	fmt.Printf("⚡ Starting Kafka consumer for topic: %s\n", nk.topic)
	ctx := context.Background()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case msg := <-nk.partitionConsumer.Messages():
			fmt.Printf("Received: %s\n", string(msg.Value))
			nk.notificationService.SendNotification(ctx, string(msg.Value))
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: shutting down\n", sig)
			run = false
		}
	}

	nk.partitionConsumer.Close()
	nk.consumer.Close()
}
