package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	_ "proxy/docs"
	dto "proxy/internal/domain/proxy"
	"proxy/internal/service/proxy"
	"proxy/pkg/server/response"
)

type ProxyHandler struct {
	proxyService *proxy.Service
}

func NewProxyHandler(s *proxy.Service) *ProxyHandler {
	return &ProxyHandler{proxyService: s}
}

// @Summary     Health check
// @Description Healthcheking.
// @Tags        health
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]string
// @Router      /health [get]
func (h *ProxyHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"Health": "OK!",
	}

	render.JSON(w, r, response)
}

func (h *ProxyHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", h.get)
	r.Post("/", h.add)

	return r
}

// @Summary	get request by id
// @Description This endpoint retrieves a request by its ID.
// @Description The ID must be provided as a path parameter.
// @Tags		requests
// @Accept		json
// @Produce	json
// @Param       id path string true "Request ID"
// @Success	200			{array}		proxy.Response
// @Failure	500			{object}	string
// @Router		/requests/{id} 	[get]
func (h *ProxyHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.proxyService.GetResponse(r.Context(), id)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, resp)
}

// @Summary     Create request
// @Description This endpoint creates a new request based on the provided payload.
// @Description The request method, URL, and headers must be specified in the request body.
// @Tags        requests
// @Accept      json
// @Produce     json
// @Param       request body dto.Request true "Request payload"
// @Success     200 {object} proxy.Response
// @Failure     500 {object} string
// @Router      /requests [post]
func (h *ProxyHandler) add(w http.ResponseWriter, r *http.Request) {
	req := dto.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	resp, err := h.proxyService.CreateRequest(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, resp)
}
