package usecase

import "github.com/8soat-grupo35/fastfood-order-production/internal/entities"

//go:generate mockgen -source=production_order.go -destination=mock/production_order.go
type ProductionOrderUseCases interface {
	SendOrderToProduction(orderId uint32) (*entities.ProductionOrder, error)
	UpdateProductionOrderStatus(orderId uint32, status string) (*entities.ProductionOrder, error)
	GetProductionOrderQueue() (*entities.ProductionOrderQueue, error)
}
