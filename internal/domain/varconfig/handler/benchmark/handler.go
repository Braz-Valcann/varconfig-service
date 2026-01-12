package benchmark

import (
	"context"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig"
	"github.com/Wizzi-Cloud/restwrapper"
	ginHandler "github.com/Wizzi-Cloud/restwrapper/gin"
	eve "github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
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
		// Verifica se existe o parâmetro "id" na URL para distinguir entre Get e List
		id, _ := w.RequestWrapper.GetPathParam("id")
		if id != nil && *id != "" {
			h.Get(w)
		} else {
			h.List(w)
		}
		return
	}
}

// Handle pra usar com Gin
func (h *Handler) HandleGin(c *gin.Context) {
	var res eve.APIGatewayProxyResponse

	wrapper := ginHandler.NewGinWrapper(c, &res)
	if wrapper == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "wrapper is nil"})
		return
	}

	// Executa o handler "core"
	h.Handle(wrapper)

	// Aplica headers simples
	for key, value := range res.Headers {
		c.Header(key, value)
	}

	// Aplica headers multi-valor (se existir)
	if len(res.MultiValueHeaders) > 0 {
		for key, values := range res.MultiValueHeaders {
			for _, v := range values {
				c.Writer.Header().Add(key, v)
			}
		}
	}

	// Define status code
	status := res.StatusCode
	if status == 0 {
		status = http.StatusOK
	}

	// Para 204/304, não envie corpo
	if status == http.StatusNoContent || status == http.StatusNotModified {
		c.Status(status)
		return
	}

	// Define Content-Type
	ct := c.Writer.Header().Get("Content-Type")
	if ct == "" {
		ct = "application/json"
	}

	// Processa body
	var payload []byte
	if res.IsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(res.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid base64 body"})
			return
		}
		payload = decoded
	} else {
		payload = []byte(res.Body)
	}

	// Escreve a resposta crua, sem re-serializar body string JSON.
	c.Data(status, ct, payload)
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
