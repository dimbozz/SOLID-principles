package service

import (
    "fmt"
    
    "SOLID-principles/internal/order"
    "SOLID-principles/internal/repository"
)

type OrderService struct {
    repo     order.RepositoryWriter
    notifier order.Notifier
}

func NewOrderService(repo order.RepositoryWriter, notifier order.Notifier) *OrderService {
    return &OrderService{
        repo:     repo,
        notifier: notifier,
    }
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
    order := &order.Order{
        Customer: customer,
        Products: fmt.Sprintf("%v", products),
        Total:    total,
        Status:   "pending",
    }
    
    if err := s.repo.CreateOrder(order); err != nil {
        return err
    }
    
    s.notifier.Send(customer)
    return nil
}
