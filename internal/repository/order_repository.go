// internal/repository/order_repository.go
package repository

import (
	"database/sql"

	"SOLID-principles/internal/order"
)

// Создаем структуру SQLiteRepo
type SQLiteRepo struct {
	db *sql.DB
}

// Конструктор для SQLiteRepo
func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

// Реализация метода CreateOrder интерфейса RepositoryWriter
func (r *SQLiteRepo) CreateOrder(order *order.Order) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, order.Products, order.Total, order.Status,
	)
	return err
}
