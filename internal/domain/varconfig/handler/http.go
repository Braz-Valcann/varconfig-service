package handler

import (
	"net/http"
	"strconv"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/dto"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: *service}
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orgID := c.Param("orgId")
	benchmarkID := c.Param("benchmark_id")

	varConfig, err := h.service.Create(c.Request.Context(), orgID, benchmarkID, req.Payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, varConfig)

}

func (h *Handler) Get(c *gin.Context) {
	idParam := c.Param("id")
	orgIdParam := c.Param("orgId")
	benchIdParam := c.Param("benchmark_id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	varConfig, err := h.service.Get(c.Request.Context(), orgIdParam, benchIdParam, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if varConfig == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Varconfig não existe!"})
		return
	}

	c.JSON(http.StatusOK, varConfig)
}

func (h *Handler) List(c *gin.Context) {
	orgIdParam := c.Param("orgId")
	benchIdParam := c.Param("benchmark_id")

	varConfigs, err := h.service.List(c.Request.Context(), orgIdParam, benchIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, varConfigs)
}

func (h *Handler) Update(c *gin.Context) {
	var req dto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orgID := c.Param("orgId")
	benchmarkID := c.Param("benchmark_id")

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	vc, err := h.service.Update(
		c.Request.Context(),
		orgID,
		benchmarkID,
		id,
		req.Payload,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vc)
}

func (h *Handler) Delete(c *gin.Context) {
	orgIdParam := c.Param("orgId")
	benchIdParam := c.Param("benchmark_id")
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Delete(c.Request.Context(), orgIdParam, benchIdParam, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
