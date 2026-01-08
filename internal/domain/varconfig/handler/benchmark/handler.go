package benchmark

/*
import (
	"net/http"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig"
	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig/service"
	"github.com/Wizzi-Cloud/restwrapper"
)

type Handler struct {
	service varconfig.IService
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(w *restwrapper.Wrapper) {
	switch w.RequestWrapper.Method() {

	case http.MethodGet:
		benchId, _ := w.RequestWrapper.GetPathParam("benchmarkId")
		if benchId != nil {
			h.Get(w)
		} else {
			h.List(w)
		}
		return
	}
}
*/
