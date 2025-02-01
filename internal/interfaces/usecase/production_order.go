package usecase

import "github.com/8soat-grupo35/fastfood-order-production/internal/entities"

type ProductionOrderUseCases interface {
	SendOrderToProduction(orderId uint32) (*entities.ProductionOrder, error)
	UpdateProductionOrderStatus(orderId uint32, status string) (*entities.ProductionOrder, error)
	GetProductionOrderQueue() (*entities.ProductionOrderQueue, error)
}
