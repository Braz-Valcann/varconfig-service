package routes

import (
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/handler"
	restHandler "github.com/Wizzi-Cloud/restwrapper/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, handler *handler.Handler, restHandler *restHandler.Handler) {
	group := r.Group("/org/:orgId/compliance/variables/:benchmark_id")
	{
		group.POST("/", handler.Create)
		// aqui está outra correção de pra poder usar o HandleGin não do arquivo de handler como estava anteriormente, mas sim do restwrapper
		group.GET("/:id", restHandler.HandleGin)
		group.GET("/", restHandler.HandleGin)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)

	}
}
