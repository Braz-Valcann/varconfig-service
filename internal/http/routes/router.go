package routes

import (
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler/benchmark"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, handler *handler.Handler, benchHandler *benchmark.Handler) {
	group := r.Group("/org/:orgId/compliance/variables/:benchmark_id")
	{
		group.POST("/", handler.Create)
		group.GET("/:id", benchHandler.HandleGin)
		group.GET("/", benchHandler.HandleGin)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)

	}
}
