package benchmark

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig"
	"github.com/Wizzi-Cloud/restwrapper"
)

type Handler struct {
	service varconfig.IService
}

func NewHandler(service varconfig.IService) *Handler {
	return &Handler{service: service}
}

// Handle implementa a interface IRestHandler do restwrapper
func (h *Handler) Handle(w *restwrapper.Wrapper) {
	switch w.RequestWrapper.Method() {
	case http.MethodGet:
		id, _ := w.RequestWrapper.GetPathParam("id")

		if id != nil && *id != "" {
			h.Get(w)
		} else {
			h.List(w)
		}
		return

	}
}

/*
	func (h *Handler) Create(w *restwrapper.Wrapper) {
		var req dto.CreateRequest

		if err := w.ShouldBindJSON(&req); err != nil {
			w.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orgID := w.Param("orgId")
		benchmarkID := w.Param("benchmark_id")

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
*/
func (h *Handler) Get(w *restwrapper.Wrapper) {
	// Extrai path parameters manualmente usando GetPathParam, eu não consegui usar o
	orgID, _ := w.RequestWrapper.GetPathParam("orgId")
	benchmarkID, _ := w.RequestWrapper.GetPathParam("benchmark_id")
	idStr, _ := w.RequestWrapper.GetPathParam("id")

	// Converte ID string para int64
	var id int64
	if idStr != nil {
		id, _ = strconv.ParseInt(*idStr, 10, 64)
	}

	// Monta o PathParameter
	pathParams := PathParameter{
		ID:          id,
		OrgID:       derefString(orgID),
		BenchmarkID: derefString(benchmarkID),
		Context:     context.Background(),
	}

	// Valida
	if err := pathParams.Validate(); err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.Get(pathParams.Context, pathParams.OrgID, pathParams.BenchmarkID, pathParams.ID)
	if err != nil {
		w.ResponseWrapper.WriteServerErrorResponse()
		return
	}

	w.ResponseWrapper.WriteSuccessResponse(http.StatusOK, response)
}

// derefString retorna o valor do ponteiro ou string vazia
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (h *Handler) List(w *restwrapper.Wrapper) {
	// Extrai org e benchmark dos path params
	orgID, _ := w.RequestWrapper.GetPathParam("orgId")
	benchmarkID, _ := w.RequestWrapper.GetPathParam("benchmark_id")

	pathParams := PathParameter{
		OrgID:       derefString(orgID),
		BenchmarkID: derefString(benchmarkID),
		Context:     context.Background(),
	}

	if err := pathParams.Validate(); err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.List(pathParams.Context, pathParams.OrgID, pathParams.BenchmarkID)
	if err != nil {
		w.ResponseWrapper.WriteServerErrorResponse()
		return
	}

	w.ResponseWrapper.WriteSuccessResponse(http.StatusOK, response)
}
