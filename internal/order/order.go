// internal/order/order.go
package order

type Order struct {
	ID       int
	Customer string
	Products string
	Total    float64
	Status   string
}

// Интерфейс RepositoryWriter - контракт для записи заказов в репозиторий
// Реализация принципа Dependency Inversion
type RepositoryWriter interface {
	CreateOrder(order *Order) error
}

// Интерфейс Notifier - контракт для отправки уведомлений
// Реализация принципа Dependency Inversion
type Notifier interface {
	Send(customer string)
}
