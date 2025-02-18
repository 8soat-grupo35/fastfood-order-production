package entities

import (
	"sort"
)

type ProductionOrderQueue struct {
	Orders []ProductionOrder
}

func (p *ProductionOrderQueue) RemoveFinishedOrders() {
	orders := []ProductionOrder{}
	for _, order := range p.Orders {
		if order.Status != FINISHED_STATUS {
			orders = append(orders, order)
		}
	}

	p.Orders = orders
}

func (p *ProductionOrderQueue) Sort() {
	sort.Slice(p.Orders, func(i, j int) bool {
		switch p.Orders[i].Status {
		case DONE_STATUS:
			return true
		case RECEIVED_STATUS:
			if p.Orders[j].Status == DONE_STATUS {
				return false
			}

			if p.Orders[j].Status == IN_PREPARATION_STATUS {
				return false
			}

			return true
		case IN_PREPARATION_STATUS:
			if p.Orders[j].Status == DONE_STATUS {
				return false
			}

			return true
		}

		return true
	})
}
