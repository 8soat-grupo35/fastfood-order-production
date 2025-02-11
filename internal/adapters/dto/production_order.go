package dto

type SendOrderToProductionDto struct {
	OrderId uint32 `json:"order_id" validate:"required"`
}

type UpdateProductionOrderStatus struct {
	Status string `json:"status"  validate:"required"`
}
