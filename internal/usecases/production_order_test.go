package usecases

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	mock_repository "github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductionOrderQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionQueue := []entities.ProductionOrder{
		{
			OrderId: 1,
			Status:  entities.RECEIVED_STATUS,
		},
		{
			OrderId: 2,
			Status:  entities.RECEIVED_STATUS,
		},
	}

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetAll().Return(productionQueue, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)

	queue, err := prodOrderUseCase.GetProductionOrderQueue()

	assert.NoError(t, err)
	assert.Equal(t, productionQueue, queue.Orders)
}

func TestGetProductionOrderQueueError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockErr := errors.New("mock error")

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetAll().Return([]entities.ProductionOrder{}, mockErr).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)

	queue, err := prodOrderUseCase.GetProductionOrderQueue()

	assert.Nil(t, queue)
	assert.EqualError(t, err, mockErr.Error())
}

func TestSendOrderToProduction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.RECEIVED_STATUS,
	}
	orderID := uint32(1)

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(orderID).Return(nil, nil).Times(1)
	mockRepo.EXPECT().Create(productionOrder).Return(&productionOrder, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)

	sendOrder, err := prodOrderUseCase.SendOrderToProduction(orderID)

	assert.NoError(t, err)
	assert.Equal(t, &productionOrder, sendOrder)
}

func TestSendOrderToProductionGetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderID := uint32(1)

	mockErr := errors.New("mock error")
	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(orderID).Return(nil, mockErr).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	sendOrder, err := prodOrderUseCase.SendOrderToProduction(orderID)

	assert.EqualError(t, err, mockErr.Error())
	assert.Nil(t, sendOrder)
}

func TestSendOrderToProductionValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderID := uint32(0)

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(orderID).Return(nil, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)

	sendOrder, err := prodOrderUseCase.SendOrderToProduction(orderID)

	assert.EqualError(t, err, "OrderId: cannot be blank.")
	assert.Nil(t, sendOrder)
}

func TestSendOrderToProductionAlreadySendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.RECEIVED_STATUS,
	}
	orderID := uint32(1)

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(orderID).Return(&productionOrder, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	sendOrder, err := prodOrderUseCase.SendOrderToProduction(orderID)

	assert.EqualError(t, err, "order already sended to production queue")
	assert.Nil(t, sendOrder)
}

func TestSendOrderToProductionCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.RECEIVED_STATUS,
	}
	orderID := uint32(1)

	mockCreateError := errors.New("mock create error")
	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(orderID).Return(nil, nil).Times(1)
	mockRepo.EXPECT().Create(productionOrder).Return(nil, mockCreateError).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)

	sendOrder, err := prodOrderUseCase.SendOrderToProduction(orderID)

	assert.EqualError(t, err, mockCreateError.Error())
	assert.Nil(t, sendOrder)
}

func TestUpdateProductionOrderStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.IN_PREPARATION_STATUS,
	}

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(productionOrder.OrderId).Return(&productionOrder, nil).Times(1)
	mockRepo.EXPECT().Update(productionOrder).Return(&productionOrder, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	updatedOrder, err := prodOrderUseCase.UpdateProductionOrderStatus(productionOrder.OrderId, productionOrder.Status)

	assert.NoError(t, err)
	assert.Equal(t, &productionOrder, updatedOrder)
}

func TestUpdateProductionOrderStatusGetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.IN_PREPARATION_STATUS,
	}

	mockGetError := errors.New("mock get error")
	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(productionOrder.OrderId).Return(nil, mockGetError).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	updatedOrder, err := prodOrderUseCase.UpdateProductionOrderStatus(productionOrder.OrderId, productionOrder.Status)

	assert.EqualError(t, err, mockGetError.Error())
	assert.Nil(t, updatedOrder)
}

func TestUpdateProductionOrderStatusNotFoundtError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.IN_PREPARATION_STATUS,
	}

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(productionOrder.OrderId).Return(nil, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	updatedOrder, err := prodOrderUseCase.UpdateProductionOrderStatus(productionOrder.OrderId, productionOrder.Status)

	assert.EqualError(t, err, "Cant find production order")
	assert.Nil(t, updatedOrder)
}

func TestUpdateProductionOrderStatusValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  "status invalid",
	}

	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(productionOrder.OrderId).Return(&productionOrder, nil).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	updatedOrder, err := prodOrderUseCase.UpdateProductionOrderStatus(productionOrder.OrderId, productionOrder.Status)

	assert.EqualError(t, err, "Status: must be between RECEBIDO, EM_PREPARACAO, PRONTO or FINALIZADO.")
	assert.Nil(t, updatedOrder)
}

func TestUpdateProductionOrderStatusUpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productionOrder := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.IN_PREPARATION_STATUS,
	}

	mockUpdateError := errors.New("mock update error")
	mockRepo := mock_repository.NewMockProductionOrderRepository(ctrl)
	mockRepo.EXPECT().GetByOrderId(productionOrder.OrderId).Return(&productionOrder, nil).Times(1)
	mockRepo.EXPECT().Update(productionOrder).Return(nil, mockUpdateError).Times(1)

	prodOrderUseCase := NewProductionOrderUseCase(mockRepo)
	updatedOrder, err := prodOrderUseCase.UpdateProductionOrderStatus(productionOrder.OrderId, productionOrder.Status)

	assert.EqualError(t, err, mockUpdateError.Error())
	assert.Nil(t, updatedOrder)
}
