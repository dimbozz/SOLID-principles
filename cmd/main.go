package main

import (
	"database/sql"
	"log"

	"SOLID-principles/internal/notification"
	"SOLID-principles/internal/repository"
	"SOLID-principles/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Dependency Injection через интерфейсы
	repo := repository.NewSQLiteRepo(db)
	emailNotifier := notification.NewEmailSender()
	// smsNotifier := notification.NewSMSSender() // Легко заменить

	// OrderService получает ИНТЕРФЕЙСЫ, а не конкретные реализации
	orderService := service.NewOrderService(repo, emailNotifier)

	err = orderService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}
}
