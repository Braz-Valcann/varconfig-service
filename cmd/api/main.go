package main

import (
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler/benchmark"
	"github.com/Braz-Valcann/varconfig-service/internal/http/routes"
	restHandler "github.com/Wizzi-Cloud/restwrapper/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	initializer := benchmark.NewHandlerInitializer().Init()
	appHandler := restHandler.NewHandler(initializer)

	r := gin.Default()

	// new benchmarkHandler
	routes.Register(r, appHandler)

	r.Run(":8080")

}
