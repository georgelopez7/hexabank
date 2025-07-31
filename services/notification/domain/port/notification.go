package port

import "context"

type NotificationService interface {
	SendNotification(ctx context.Context, message string) error
}

type DiscordClient interface {
	SendMessage(message string) error
}
