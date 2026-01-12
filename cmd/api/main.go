package main

import (
	"context"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler/benchmark"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/service"
	dynamodbConfig "github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/storage/dynamodb"
	"github.com/Braz-Valcann/varconfig-service/internal/http/routes"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, _ := config.LoadDefaultConfig(context.TODO())

	ddb := dynamodb.NewFromConfig(cfg)

	repo := dynamodbConfig.New(ddb, "var_configs")
	service := service.New(repo)
	handler := handler.New(service)
	// new benchmarkHandler := benchmark.NewHandler(service)
	benchmarkHandler := benchmark.NewHandler(service)

	r := gin.Default()

	// new benchmarkHandler
	routes.Register(r, handler, benchmarkHandler)

	r.Run(":8080")

}
