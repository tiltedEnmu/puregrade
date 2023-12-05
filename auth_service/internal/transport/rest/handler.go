package rest

import (
	"net/http"

	"github.com/puregrade/puregrade-auth/internal/service"
)

type Handler struct {
	services        service.Service
	UserServiceAddr string
	Routes          *http.ServeMux
}

func NewHandler(services service.Service, userServiceAddr string) *Handler {
	return &Handler{
		services:        services,
		UserServiceAddr: userServiceAddr,
		Routes:          http.NewServeMux(),
	}
}

func (h *Handler) InitRoutes() {
	h.Routes.HandleFunc("/sign-up", h.signUp)
	h.Routes.HandleFunc("/sign-in", h.signIn)
	h.Routes.HandleFunc("/refresh", h.refresh)
}
