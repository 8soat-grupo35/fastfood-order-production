package handlers

import (
	"net/http"
	"strconv"

	"github.com/8soat-grupo35/fastfood-order-production/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/usecase"
	"github.com/labstack/echo/v4"
)

type ProductionOrderHandler struct {
	productionOrderUseCases usecase.ProductionOrderUseCases
}

func NewProductionOrderHandler(usecase usecase.ProductionOrderUseCases) ProductionOrderHandler {

	return ProductionOrderHandler{
		productionOrderUseCases: usecase,
	}
}

func (h *ProductionOrderHandler) SendOrderToProduction(echo echo.Context) error {
	sendOrderToProductionDto := dto.SendOrderToProductionDto{}

	err := echo.Bind(&sendOrderToProductionDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = echo.Validate(sendOrderToProductionDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	orderSend, err := h.productionOrderUseCases.SendOrderToProduction(sendOrderToProductionDto.OrderId)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, orderSend)
}

func (h *ProductionOrderHandler) UpdateProductionOrderStatus(echo echo.Context) error {
	updateProductionOrderStatusDto := dto.UpdateProductionOrderStatus{}
	orderId, err := strconv.Atoi(echo.Param("orderId"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = echo.Bind(&updateProductionOrderStatusDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = echo.Validate(updateProductionOrderStatusDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	productionOrderUpdated, err := h.productionOrderUseCases.UpdateProductionOrderStatus(uint32(orderId), updateProductionOrderStatusDto.Status)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, productionOrderUpdated)
}

func (h *ProductionOrderHandler) GetProductionOrderQueue(echo echo.Context) error {
	productionOrderQueue, err := h.productionOrderUseCases.GetProductionOrderQueue()

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, productionOrderQueue.Orders)
}
