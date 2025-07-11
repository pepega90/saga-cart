package contracts

import "order_service_saga/internal/domain"

type OrderRepoContract interface {
	CreateOrder(domain.Order) (*domain.Order, error)
	GerOrders() (*[]domain.Order, error)
	UpdateOrderStatus(int, bool) error
}

type OrderServiceContract interface {
	CreateOrder(domain.Order) (*domain.Order, error)
	GerOrders() (*[]domain.Order, error)
}
