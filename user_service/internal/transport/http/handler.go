package rest

import (
	"net/http"

	"github.com/tiltedEnmu/puregrade_user/internal/service"
)

type Handler struct {
	services service.Service
	ServeMux *http.ServeMux
}

func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
		ServeMux: http.NewServeMux(),
	}
}

func (h *Handler) InitRoutes() {
	h.ServeMux.HandleFunc("/sing-up", h.singUp)
	h.ServeMux.HandleFunc("/sing-in", h.singIn)
	h.ServeMux.HandleFunc("/refresh", h.refresh)
}
