// internal/notification/notification_service.go
package notification

// NotificationService — только отправка уведомлений
type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) SendEmailNotification (customer string) {
	println("Уведомление отправлено клиенту", customer)
}
