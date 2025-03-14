package gateways

import (
	"fmt"
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

	for _, item := range value {
		orders = append(orders, entities.ProductionOrder{
			OrderId: uint32(item["ID"].(float64)),
			Status:  item["Status"].(string),
		})
	}

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

	order = p.convertDynamoToEntity(value)

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
	fmt.Println(order)
	value, err := p.dynamo.UpdateValue("ID", order.OrderId, "Status", order.Status)

	if err != nil {
		return nil, err
	}

	updatedProductionOrder = p.convertDynamoToEntity(value)

	return updatedProductionOrder, nil
}

func (p productionOrderGateway) convertDynamoToEntity(item map[string]interface{}) *entities.ProductionOrder {
	return &entities.ProductionOrder{
		OrderId: uint32(item["ID"].(float64)),
		Status:  item["Status"].(string),
	}
}

func NewProductionOrderGateway(orm external.DynamoAdapter) repository.ProductionOrderRepository {
	orm.SetTable("production_order")
	return &productionOrderGateway{
		dynamo: orm,
	}
}
