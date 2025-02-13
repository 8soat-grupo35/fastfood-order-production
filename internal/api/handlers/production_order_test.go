package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/8soat-grupo35/fastfood-order-production/external"
	"github.com/8soat-grupo35/fastfood-order-production/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	mock_usecase "github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/usecase/mock"
	"github.com/8soat-grupo35/fastfood-order-production/internal/interfaces/utils"
	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var echoContext = func(method string, targetPath string, body io.Reader) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &external.HandlerCustomValidator{
		Validator: validator.New(),
	}
	req := httptest.NewRequest(method, targetPath, body)
	if body != nil {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), req, rec
}

func TestProductionOrderHandler_GetProductionOrderQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	useCase := mock_usecase.NewMockProductionOrderUseCases(ctrl)

	queue := entities.ProductionOrderQueue{
		Orders: []entities.ProductionOrder{
			{
				OrderId: 1,
				Status:  "RECEBIDO",
			},
		},
	}

	testCases := []utils.TestCase{
		{
			Name: "Should return production order queue successfully",
			SetupMocks: func() interface{} {
				useCase.EXPECT().GetProductionOrderQueue().Return(&queue, nil).Times(1)
				res, err := json.Marshal(queue.Orders)
				assert.NoError(t, err)
				return map[string]interface{}{
					"code": http.StatusOK,
					"body": string(res),
				}
			},
			WantErr: false,
		},
		{
			Name: "Should return 500 when cant find order queue successfully",
			SetupMocks: func() interface{} {
				mockErr := errors.New("mock error")
				useCase.EXPECT().GetProductionOrderQueue().Return(nil, mockErr).Times(1)
				res, err := json.Marshal(mockErr.Error())
				assert.NoError(t, err)
				return map[string]interface{}{
					"code": http.StatusInternalServerError,
					"body": string(res),
				}
			},
			WantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()
			ctx, _, res := echoContext(http.MethodGet, "/production/queue", nil)
			handler := NewProductionOrderHandler(useCase)
			err := handler.GetProductionOrderQueue(ctx)

			responseBody := strings.ReplaceAll(res.Body.String(), "\n", "")

			assert.Equal(t, expectedValue, map[string]interface{}{
				"code": res.Code,
				"body": responseBody,
			})

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}

}

func TestProductionOrderHandler_SendOrderToProduction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	useCase := mock_usecase.NewMockProductionOrderUseCases(ctrl)

	sendOrderDto := dto.SendOrderToProductionDto{
		OrderId: 1,
	}

	sendOrderEntity := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.RECEIVED_STATUS,
	}

	sendOrderDtoStr, err := json.Marshal(sendOrderDto)
	assert.NoError(t, err)

	testCases := []utils.TestCase{
		{
			Name: "Should send order to production queue successfully",
			SetupMocks: func() interface{} {
				useCase.EXPECT().SendOrderToProduction(sendOrderDto.OrderId).Return(&sendOrderEntity, nil).Times(1)
				res, err := json.Marshal(sendOrderEntity)
				assert.NoError(t, err)
				return map[string]interface{}{
					"code": http.StatusOK,
					"body": string(res),
				}
			},
			WantErr: false,
		},
		{
			Name: "Should return 500 when cant send order to production queue successfully",
			SetupMocks: func() interface{} {
				mockErr := errors.New("mock error")
				useCase.EXPECT().SendOrderToProduction(sendOrderDto.OrderId).Return(nil, mockErr).Times(1)
				res, err := json.Marshal(mockErr.Error())
				assert.NoError(t, err)
				return map[string]interface{}{
					"code": http.StatusInternalServerError,
					"body": string(res),
				}
			},
			WantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()
			ctx, _, res := echoContext(http.MethodGet, "/production/order/send", strings.NewReader(string(sendOrderDtoStr)))
			handler := NewProductionOrderHandler(useCase)
			err := handler.SendOrderToProduction(ctx)

			responseBody := strings.ReplaceAll(res.Body.String(), "\n", "")

			assert.Equal(t, expectedValue, map[string]interface{}{
				"code": res.Code,
				"body": responseBody,
			})

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestProductionOrderHandler_UpdateProductionOrderStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	useCase := mock_usecase.NewMockProductionOrderUseCases(ctrl)

	updateOrderDto := dto.UpdateProductionOrderStatus{
		Status: entities.DONE_STATUS,
	}

	updateOrderEntity := entities.ProductionOrder{
		OrderId: 1,
		Status:  entities.DONE_STATUS,
	}

	updateOrderDtoStr, err := json.Marshal(updateOrderDto)
	assert.NoError(t, err)

	testCases := []utils.TestCase{
		{
			Name: "Should update order on production queue successfully",
			SetupMocks: func() interface{} {
				useCase.EXPECT().UpdateProductionOrderStatus(updateOrderEntity.OrderId, updateOrderDto.Status).Return(&updateOrderEntity, nil).Times(1)
				res, err := json.Marshal(updateOrderEntity)
				assert.NoError(t, err)
				return map[string]interface{}{
					"code": http.StatusOK,
					"body": string(res),
				}
			},
			WantErr: false,
		},
		{
			Name: "Should return 500 when cant send order to production queue successfully",
			SetupMocks: func() interface{} {
				mockErr := errors.New("mock error")
				useCase.EXPECT().UpdateProductionOrderStatus(updateOrderEntity.OrderId, updateOrderDto.Status).Return(nil, mockErr).Times(1)
				res, err := json.Marshal(mockErr.Error())
				assert.NoError(t, err)
				return map[string]interface{}{
					"code": http.StatusInternalServerError,
					"body": string(res),
				}
			},
			WantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			expectedValue := tt.SetupMocks()
			ctx, _, res := echoContext(http.MethodGet, "/production/order/1/status", strings.NewReader(string(updateOrderDtoStr)))
			ctx.SetParamNames("orderId")
			ctx.SetParamValues("1")

			handler := NewProductionOrderHandler(useCase)
			err := handler.UpdateProductionOrderStatus(ctx)

			responseBody := strings.ReplaceAll(res.Body.String(), "\n", "")

			assert.Equal(t, expectedValue, map[string]interface{}{
				"code": res.Code,
				"body": responseBody,
			})

			if tt.WantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestProductionOrderHandler_UpdateProductionOrderStatus_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	useCase := mock_usecase.NewMockProductionOrderUseCases(ctrl)

	ctx, _, res := echoContext(http.MethodGet, "/production/order/1/status", nil)

	handler := NewProductionOrderHandler(useCase)
	err := handler.UpdateProductionOrderStatus(ctx)

	assert.NoError(t, err)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}
