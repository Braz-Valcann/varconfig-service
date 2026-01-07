package routes

import (
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, handler *handler.Handler) {
	group := r.Group("/org/:orgId/compliance/variables/:benchmark_id")
	{
		group.POST("/", handler.Create)
		group.GET("/:id", handler.Get)
		group.GET("/", handler.List)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)

	}
}
