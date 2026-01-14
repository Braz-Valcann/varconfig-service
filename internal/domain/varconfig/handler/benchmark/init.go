package benchmark

import (
	"context"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/service"
	dynamodbConfig "github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/storage/dynamodb"
	"github.com/Wizzi-Cloud/restwrapper/handler"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type HandlerInitializer struct{}

func NewHandlerInitializer() *HandlerInitializer {
	return &HandlerInitializer{}
}

func (h *HandlerInitializer) Init() handler.IRestHandler {
	cfg, _ := config.LoadDefaultConfig(context.TODO())

	ddb := dynamodb.NewFromConfig(cfg)

	repo := dynamodbConfig.New(ddb, "var_configs")
	service := service.New(repo)

	return NewHandler(service)
}
