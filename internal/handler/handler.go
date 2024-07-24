package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"proxy/internal/handler/http"
	"proxy/internal/service/proxy"
	"proxy/pkg/server/router"
)

type Dependencies struct {
	ProxyService *proxy.Service
}

type Configuration func(h *Handler) error

type Handler struct {
	dependencies Dependencies

	HTTP *chi.Mux
}

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	h = &Handler{
		dependencies: d,
	}

	for _, cfg := range configs {
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.HTTP = router.New()

		h.HTTP.Use(middleware.Timeout(60))
		h.HTTP.Get("/swagger/*", httpSwagger.WrapHandler)

		proxyHandler := http.NewProxyHandler(h.dependencies.ProxyService)

		h.HTTP.Get("/health", proxyHandler.HealthCheck)

		h.HTTP.Route("/", func(r chi.Router) {
			r.Mount("/requests", proxyHandler.Routes())
		})

		return
	}
}
