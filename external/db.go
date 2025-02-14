package external

import (
	"context"
	"log"

	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
)

var (
	DB *dynamo.DB
)

func ConectaDB(config Config) *dynamo.DB {

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithDefaultRegion("us-east-1"))

	if err != nil {
		log.Println(err.Error())
		log.Panic("Erro na conexao com banco de dados")
	}

	if config.Environment == "development" {
		baseURL := "http://localhost:4566"
		cfg.BaseEndpoint = &baseURL
	}

	DB = dynamo.New(cfg)

	err = DB.CreateTable("production_order", entities.ProductionOrder{}).OnDemand(true).Run(context.TODO())

	if err != nil {
		log.Println(err.Error())
	}

	return DB
}
