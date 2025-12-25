// internal/service/order_service.go
package service

import (
	"fmt"

	"SOLID-principles/internal/order"
)

// OrderService зависит ТОЛЬКО от ИНТЕРФЕЙСОВ
// НЕ зависит от конкретных структур SQLiteRepo, EmailSender, SMSSender
type OrderService struct {
	repo     order.RepositoryWriter
	notifier order.Notifier
}

// Конструктор с внедрением зависимостей через интерфейсы (Dependency Injection)
func NewOrderService(repo order.RepositoryWriter, notifier order.Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

// Бизнес-логика создания заказа
func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := &order.Order{
		Customer: customer,
		Products: fmt.Sprintf("%v", products),
		Total:    total,
		Status:   "pending",
	}

	// Вызов интерфейса репозитория
	if err := s.repo.CreateOrder(order); err != nil {
		return err
	}

	// Вызов интерфейса уведомлений
	s.notifier.Send(customer)
	return nil
}
