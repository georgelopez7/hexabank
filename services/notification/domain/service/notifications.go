package service

import (
	"context"
	"hexabank/services/notification/domain/port"
)

type NotificationService struct {
	discordClient port.DiscordClient
}

func NewNotificationService(discordClient port.DiscordClient) *NotificationService {
	return &NotificationService{
		discordClient: discordClient,
	}
}

func (s *NotificationService) SendNotification(ctx context.Context, message string) error {
	err := s.discordClient.SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}
