package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"

	_ "github.com/8soat-grupo35/fastfood-order-production/docs"
	"github.com/8soat-grupo35/fastfood-order-production/external"
	"github.com/8soat-grupo35/fastfood-order-production/internal/api/handlers"
	"github.com/8soat-grupo35/fastfood-order-production/internal/gateways"
	"github.com/8soat-grupo35/fastfood-order-production/internal/usecases"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start() {
	cfg := external.GetConfig()
	fmt.Println(cfg)

	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))
	app := newApp(cfg)
	app.Logger.Fatal(app.Start(cfg.ServerHost))
}

// @title Swagger Fastfood App API
// @version 1.0
// @description This is a sample API from Fastfood App.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1
func newApp(cfg external.Config) *echo.Echo {
	database := external.ConectaDB(cfg)
	app := echo.New()
	app.Validator = &external.HandlerCustomValidator{
		Validator: validator.New(),
	}
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.GET("/", func(echo echo.Context) error {
		return echo.JSON(http.StatusOK, "Alive")
	})

	productionOrderGateway := gateways.NewProductionOrderGateway(
		external.NewDynamoAdapter(database),
	)
	productionOrderHandler := handlers.NewProductionOrderHandler(
		usecases.NewProductionOrderUseCase(
			productionOrderGateway,
		),
	)
	app.GET("/production/queue", productionOrderHandler.GetProductionOrderQueue)
	app.POST("/production/order/send", productionOrderHandler.SendOrderToProduction)
	app.PUT("/production/order/:orderId/status", productionOrderHandler.UpdateProductionOrderStatus)

	return app
}
