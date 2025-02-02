package gateways

import (
	"context"
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/repository"
	"github.com/guregu/dynamo/v2"
	"strings"
)

type productionOrderGateway struct {
	orm   *dynamo.DB
	table dynamo.Table
}

func (p productionOrderGateway) GetAll() (orders []entities.ProductionOrder, err error) {
	err = p.table.Scan().All(context.TODO(), &orders)

	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (p productionOrderGateway) GetByOrderId(orderId uint32) (order *entities.ProductionOrder, err error) {
	err = p.table.Get("ID", orderId).One(context.TODO(), &order)

	if err != nil {

		if strings.Contains(err.Error(), "no item found") {
			return nil, nil
		}

		return order, err
	}

	return order, nil
}

func (p productionOrderGateway) Create(order entities.ProductionOrder) (*entities.ProductionOrder, error) {
	err := p.table.Put(order).Run(context.TODO())

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (p productionOrderGateway) Update(order entities.ProductionOrder) (updatedProductionOrder *entities.ProductionOrder, err error) {
	err = p.table.Update("ID", order.OrderId).
		Set("Status", order.Status).
		Value(context.TODO(), &updatedProductionOrder)

	if err != nil {
		return nil, err
	}

	return updatedProductionOrder, nil
}

func NewProductionOrderGateway(orm *dynamo.DB) repository.ProductionOrderRepository {
	return &productionOrderGateway{
		orm:   orm,
		table: orm.Table("production_order"),
	}
}
