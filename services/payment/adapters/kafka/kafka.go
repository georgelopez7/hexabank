package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type NotificationProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewNotificationProducer(brokers []string, topic string) (*NotificationProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &NotificationProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *NotificationProducer) Send(message string) error {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(msg)
	return err
}

func (p *NotificationProducer) Close() error {
	return p.producer.Close()
}
