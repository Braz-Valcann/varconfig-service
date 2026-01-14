package main

import (
	"context"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler/benchmark"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/service"
	dynamodbConfig "github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/storage/dynamodb"
	"github.com/Braz-Valcann/varconfig-service/internal/http/routes"
	restHandler "github.com/Wizzi-Cloud/restwrapper/handler"
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
	// aqui eu só passei pro Handler do restwrapper pra ter acesso ao HandleGin. Uma correção a ser ver se o Dr. Matheus achar melhor fazer um init, o problema é que eu teria que criar todas injeções de dependências num arquivo separado.
	appHandler := restHandler.NewHandler(benchmarkHandler)

	r := gin.Default()

	// new benchmarkHandler
	routes.Register(r, handler, appHandler)

	r.Run(":8080")

}
