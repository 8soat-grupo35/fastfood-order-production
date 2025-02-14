package gateways

import (
	"errors"
	"testing"

	mock_external "github.com/8soat-grupo35/fastfood-order-production/external/mock"
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductionOrderGateway_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAdapter := mock_external.NewMockDynamoAdapter(ctrl)
	mockAdapter.EXPECT().SetTable(gomock.Any()).AnyTimes().Return()

	testCases := []utils.TestCase{
		{
			Name: "should return all orders",
			SetupMocks: func() interface{} {
				expectedOrders := []entities.ProductionOrder{
					{
						OrderId: 1,
						Status:  "RECEBIDO",
					},
					{
						OrderId: 2,
						Status:  "EM_PREPARACAO",
					},
				}

				response := []map[string]interface{}{
					{
						"ID":     float64(1),
						"Status": "RECEBIDO",
					},
					{
						"ID":     float64(2),
						"Status": "EM_PREPARACAO",
					},
				}

				mockAdapter.EXPECT().GetAll().Return(response, nil).Times(1)

				return expectedOrders
			},
			WantErr: false,
		},
		{
			Name: "should return error if dynamo fails",
			SetupMocks: func() interface{} {
				expectedValue := []entities.ProductionOrder{}

				mockAdapter.EXPECT().GetAll().Return([]map[string]interface{}{}, errors.New("teste")).Times(1)

				return expectedValue
			},
			WantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()

			got, err := NewProductionOrderGateway(mockAdapter).GetAll()

			assert.Equal(t, expectedValue, got)

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestProductionOrderGateway_GetByOrderId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAdapter := mock_external.NewMockDynamoAdapter(ctrl)
	mockAdapter.EXPECT().SetTable(gomock.Any()).AnyTimes().Return()

	testCases := []utils.TestCase{
		{
			Name: "should return the order successfully",
			SetupMocks: func() interface{} {
				expectedOrder := &entities.ProductionOrder{
					OrderId: 1,
					Status:  "RECEBIDO",
				}

				mockAdapter.EXPECT().GetOneByKey("ID", uint32(1)).Return(expectedOrder, nil).Times(1)

				return expectedOrder
			},
			WantErr: false,
		},
		{
			Name: "should return error if dynamo fails",
			SetupMocks: func() interface{} {
				var expectedValue *entities.ProductionOrder = nil

				mockAdapter.EXPECT().GetOneByKey("ID", uint32(1)).Return(expectedValue, errors.New("teste")).Times(1)

				return expectedValue
			},
			WantErr: true,
		},
		{
			Name: "should not return error if no item is found",
			SetupMocks: func() interface{} {
				var expectedValue *entities.ProductionOrder = nil

				mockAdapter.EXPECT().GetOneByKey("ID", uint32(1)).Return(expectedValue, errors.New("no item found")).Times(1)

				return expectedValue
			},
			WantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()

			got, err := NewProductionOrderGateway(mockAdapter).GetByOrderId(1)

			assert.Equal(t, expectedValue, got)

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestProductionOrderGateway_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAdapter := mock_external.NewMockDynamoAdapter(ctrl)
	mockAdapter.EXPECT().SetTable(gomock.Any()).AnyTimes().Return()

	orderToCreate := entities.ProductionOrder{
		OrderId: 1,
		Status:  "RECEBIDO",
	}

	testCases := []utils.TestCase{
		{
			Name: "should create the order successfully",
			SetupMocks: func() interface{} {

				mockAdapter.EXPECT().Create(orderToCreate).Return(nil).Times(1)

				return &orderToCreate
			},
			WantErr: false,
		},
		{
			Name: "should return error if dynamo fails",
			SetupMocks: func() interface{} {
				var expectedValue *entities.ProductionOrder = nil

				mockAdapter.EXPECT().Create(orderToCreate).Return(errors.New("teste")).Times(1)

				return expectedValue
			},
			WantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()

			got, err := NewProductionOrderGateway(mockAdapter).Create(orderToCreate)

			assert.Equal(t, expectedValue, got)

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestProductionOrderGateway_UpdateValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAdapter := mock_external.NewMockDynamoAdapter(ctrl)
	mockAdapter.EXPECT().SetTable(gomock.Any()).AnyTimes().Return()

	orderToUpdate := entities.ProductionOrder{
		OrderId: 1,
		Status:  "RECEBIDO",
	}

	testCases := []utils.TestCase{
		{
			Name: "should update the order successfully",
			SetupMocks: func() interface{} {

				mockAdapter.EXPECT().UpdateValue("ID", orderToUpdate.OrderId, "Status", orderToUpdate.Status).Return(&orderToUpdate, nil).Times(1)

				return &orderToUpdate
			},
			WantErr: false,
		},
		{
			Name: "should return error if dynamo fails",
			SetupMocks: func() interface{} {
				var expectedValue *entities.ProductionOrder = nil

				mockAdapter.EXPECT().UpdateValue("ID", orderToUpdate.OrderId, "Status", orderToUpdate.Status).Return(expectedValue, errors.New("teste")).Times(1)

				return expectedValue
			},
			WantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()

			got, err := NewProductionOrderGateway(mockAdapter).Update(orderToUpdate)

			assert.Equal(t, expectedValue, got)

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
