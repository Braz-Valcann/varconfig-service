package benchmark

import (
	"net/http"
	"strconv"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/dto"
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
		_, resp := w.RequestWrapper.GetPathParam("id")
		if resp {
			h.Get(w)
		} else {
			h.List(w)
		}
		return
	case http.MethodPost:
		h.Create(w)
	case http.MethodPut:
		h.Update(w)
	case http.MethodDelete:
		h.Delete(*w)
	}

}

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
	}

	// Valida
	if err := pathParams.Validate(); err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.Get(pathParams.OrgID, pathParams.BenchmarkID, pathParams.ID)
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
	}

	if err := pathParams.Validate(); err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.List(pathParams.OrgID, pathParams.BenchmarkID)
	if err != nil {
		w.ResponseWrapper.WriteServerErrorResponse()
		return
	}

	w.ResponseWrapper.WriteSuccessResponse(http.StatusOK, response)
}

// implementação tudo no restwrapper

func (h *Handler) Create(w *restwrapper.Wrapper) {
	var req dto.CreateRequest

	if err := w.RequestWrapper.BindBody(&req); err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	orgID, _ := w.RequestWrapper.GetPathParam("orgId")
	benchmarkID, _ := w.RequestWrapper.GetPathParam("benchmark_id")

	varConfig, err := h.service.Create(*orgID, *benchmarkID, req.Payload)

	if err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusInternalServerError, err.Error())
		return
	}

	w.ResponseWrapper.WriteSuccessResponse(http.StatusOK, varConfig)
}

func (h *Handler) Update(w *restwrapper.Wrapper) {
	var req dto.UpdateRequest
	if err := w.RequestWrapper.BindBody(&req); err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	orgID, _ := w.RequestWrapper.GetPathParam("orgId")
	benchmarkID, _ := w.RequestWrapper.GetPathParam("benchmark_id")

	idParam, _ := w.RequestWrapper.GetPathParam("id")
	id, err := strconv.ParseInt(*idParam, 10, 64)
	if err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	vc, err := h.service.Update(
		*orgID,
		*benchmarkID,
		id,
		req.Payload,
	)
	if err != nil {
		w.ResponseWrapper.WriteServerErrorResponse()
		return
	}

	w.ResponseWrapper.WriteSuccessResponse(http.StatusOK, vc)
}

func (h *Handler) Delete(w restwrapper.Wrapper) {
	orgIdParam, _ := w.RequestWrapper.GetPathParam("orgId")
	benchIdParam, _ := w.RequestWrapper.GetPathParam("benchmark_id")
	idParam, _ := w.RequestWrapper.GetPathParam("id")

	id, err := strconv.ParseInt(*idParam, 10, 64)
	if err != nil {
		w.ResponseWrapper.WriteClientErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(*orgIdParam, *benchIdParam, id)

	if err != nil {
		w.ResponseWrapper.WriteServerErrorResponse()
		return
	}

	w.ResponseWrapper.WriteSuccessResponse(http.StatusNoContent, "Deleted")

}
