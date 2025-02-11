package gateways

import (
	"strings"

	"github.com/8soat-grupo35/fastfood-order-production/external"
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/repository"
)

type productionOrderGateway struct {
	dynamo external.DynamoAdapter
}

func (p productionOrderGateway) GetAll() (orders []entities.ProductionOrder, err error) {
	value, err := p.dynamo.GetAll()

	if err != nil {
		return []entities.ProductionOrder{}, err
	}

	orders = value.([]entities.ProductionOrder)

	return orders, nil
}

func (p productionOrderGateway) GetByOrderId(orderId uint32) (order *entities.ProductionOrder, err error) {
	value, err := p.dynamo.GetOneByKey("ID", orderId)

	if err != nil {

		if strings.Contains(err.Error(), "no item found") {
			return nil, nil
		}

		return order, err
	}

	order = value.(*entities.ProductionOrder)

	return order, nil
}

func (p productionOrderGateway) Create(order entities.ProductionOrder) (*entities.ProductionOrder, error) {
	err := p.dynamo.Create(order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (p productionOrderGateway) Update(order entities.ProductionOrder) (updatedProductionOrder *entities.ProductionOrder, err error) {
	value, err := p.dynamo.UpdateValue("ID", order.OrderId, "Status", order.Status)

	if err != nil {
		return nil, err
	}

	updatedProductionOrder = value.(*entities.ProductionOrder)

	return updatedProductionOrder, nil
}

func NewProductionOrderGateway(orm external.DynamoAdapter) repository.ProductionOrderRepository {
	orm.SetTable("production_order")
	return &productionOrderGateway{
		dynamo: orm,
	}
}
