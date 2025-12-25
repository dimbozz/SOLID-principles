// internal/notification/notification_service.go
package notification

// EmailSender реализует интерфейс Notifier
// Метод sendEmailNotification перенесен сюда и переименован в Send()
// Структура EmailSender выделена отдельно
type EmailSender struct{}

// Конструктор EmailSender
func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

// Реализация метода Send() интерфейса Notifier
// (перенесен из OrderSystem.sendEmailNotification и переименован)
func (e *EmailSender) Send(customer string) {
	println("Email уведомление отправлено клиенту", customer)
}

// Структура SMSSender реализует интерфейс Notifier
type SMSSender struct{}

// Конструктор SMSSender
func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

// Реализация метода Send() для SMS
func (s *SMSSender) Send(customer string) {
	println("SMS уведомление отправлено клиенту", customer)
}
