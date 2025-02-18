package usecases

import (
	custom_errors "github.com/8soat-grupo35/fastfood-order-production/internal/api/errors"
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/repository"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/usecase"
)

type productionOrderService struct {
	productionOrderRepository repository.ProductionOrderRepository
}

func NewProductionOrderUseCase(productionOrderRepository repository.ProductionOrderRepository) usecase.ProductionOrderUseCases {
	return &productionOrderService{
		productionOrderRepository: productionOrderRepository,
	}
}

// GetProductionOrderQueue implements usecase.ProductionOrderUseCases.
func (p *productionOrderService) GetProductionOrderQueue() (*entities.ProductionOrderQueue, error) {
	productionOrders, err := p.productionOrderRepository.GetAll()

	if err != nil {
		return nil, &custom_errors.DatabaseError{
			Message: err.Error(),
		}
	}

	productionQueue := entities.ProductionOrderQueue{
		Orders: productionOrders,
	}

	productionQueue.RemoveFinishedOrders()
	productionQueue.Sort()

	return &productionQueue, nil
}

// SendOrderToProduction implements usecase.ProductionOrderUseCases.
func (p *productionOrderService) SendOrderToProduction(orderId uint32) (*entities.ProductionOrder, error) {

	foundProductionOrder, err := p.productionOrderRepository.GetByOrderId(orderId)

	if err != nil {
		return nil, err
	}

	if foundProductionOrder != nil {
		return nil, &custom_errors.BadRequestError{
			Message: "order already sended to production queue",
		}
	}

	productionOrder := entities.ProductionOrder{
		OrderId: orderId,
		Status:  entities.RECEIVED_STATUS,
	}

	err = productionOrder.Validate()

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	createdProductionOrder, err := p.productionOrderRepository.Create(productionOrder)

	if err != nil {
		return nil, err
	}

	return createdProductionOrder, nil
}

func (p *productionOrderService) UpdateProductionOrderStatus(orderId uint32, status string) (*entities.ProductionOrder, error) {
	foundProductionOrder, err := p.productionOrderRepository.GetByOrderId(orderId)

	if err != nil {
		return nil, err
	}

	if foundProductionOrder == nil {
		return nil, &custom_errors.BadRequestError{
			Message: "Cant find production order",
		}
	}

	foundProductionOrder.Status = status

	err = foundProductionOrder.Validate()

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	updatedProductionOrder, err := p.productionOrderRepository.Update(*foundProductionOrder)

	if err != nil {
		return nil, &custom_errors.DatabaseError{
			Message: err.Error(),
		}
	}

	return updatedProductionOrder, nil
}
