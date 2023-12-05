// Package rest
package rest

import (
	"net/http"

	"github.com/tiltedEnmu/puregrade_post/internal/service"
)

type Handler struct {
	services service.Service
	Routes   *http.ServeMux
}

// ! NewHandler
func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
		Routes:   http.NewServeMux(),
	}
}

func (h *Handler) InitRoutes() {
	h.Routes.HandleFunc("/create", h.create)
	h.Routes.HandleFunc("/get", h.get)
	h.Routes.HandleFunc("/delete", h.delete)
}
