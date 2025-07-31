package port

type NotificationProducer interface {
	Send(message string) error
	Close() error
}
