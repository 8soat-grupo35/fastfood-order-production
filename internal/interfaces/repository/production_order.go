package repository

import "github.com/8soat-grupo35/fastfood-order-production/internal/entities"

//go:generate mockgen -source=production_order.go -destination=mock/production_order.go
type ProductionOrderRepository interface {
	GetAll() ([]entities.ProductionOrder, error)
	GetByOrderId(orderId uint32) (*entities.ProductionOrder, error)
	Create(order entities.ProductionOrder) (*entities.ProductionOrder, error)
	Update(order entities.ProductionOrder) (*entities.ProductionOrder, error)
}
