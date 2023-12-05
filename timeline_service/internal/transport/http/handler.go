// Package rest
package rest

import (
	"net/http"

	"github.com/tiltedEnmu/puregrade_timeline/internal/service"
)

type Handler struct {
	services service.Service
	Routes   *http.ServeMux
}

func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
		Routes:   http.NewServeMux(),
	}
}

func (h *Handler) InitRoutes() {
	h.Routes.HandleFunc("/get", h.getAll)
	h.Routes.HandleFunc("/get-latest", h.getLatest)
}
