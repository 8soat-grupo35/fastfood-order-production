package repository

import "github.com/8soat-grupo35/fastfood-order-production/internal/entities"

type ProductionOrderRepository interface {
	GetAll() ([]entities.ProductionOrder, error)
	GetByOrderId(orderId uint32) (*entities.ProductionOrder, error)
	Create(order entities.ProductionOrder) (*entities.ProductionOrder, error)
	Update(order entities.ProductionOrder) (*entities.ProductionOrder, error)
}
