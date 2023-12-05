package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tiltedEnmu/puregrade_user/internal/entities"
)

type singInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input entities.User

	buf := make([]byte, 256)
	n, err := r.Body.Read(buf)

	// erase empty bytes
	if err == io.EOF {
		buf = buf[:n]
	}

	if err := json.Unmarshal(buf, &input); err != nil {
		log.Println("Unmarshal failed: /sing-up ", err.Error())
		_, _ = w.Write([]byte("Unmarshal failed!"))
		return
	}

	if b, err := json.Marshal(input); err == nil {
		log.Println("/sing-up | ", string(b))
	}

	id, err := h.services.CreateUser(input) // call to UserService

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte("User registration failed"))
		return
	}

	access, refresh, err := h.services.GenerateTokens(id)
	if err != nil {
		log.Println("Generate tokens failed: /sing-up ", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Generate tokens failed"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("{\"access_token\":\"%s\",\"refresh_token\":\"%s\"}", access, refresh)))
}

func (h *Handler) singIn(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Close request body error")
		}
	}(r.Body)

	var userId int64
	var input singInDTO

	buf, err := io.ReadAll(r.Body);

	if err := json.Unmarshal(buf, &input); err != nil {
		log.Println("Unmarshal failed: /sing-in ", err.Error())
		_, _ = w.Write([]byte("Unmarshal failed"))
		return
	}

	userId = 123 // call to UserService & return userId (if username&password is valid)

	access, refresh, err := h.services.GenerateTokens(userId)
	if err != nil {
		log.Println("Generate tokens failed: /sing-in ", err.Error())
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Generate tokens failed"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("{\"access_token\":\"%s\",\"refresh_token\":\"%s\"}", access, refresh)))
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	refresh := r.URL.Query().Get("refresh")

	id, err := h.services.GetUserId(refresh)
	if err != nil {
		log.Error("Searching token failed: /refresh ", err.Error())
		w.WriteHeader(404)
		w.Write([]byte("refresh token was not found"))
		return
	}

	access, refresh, err := h.services.GenerateTokens(id)
	if err != nil {
		log.Error("Generate tokens failed: /refresh ", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Generate tokens failed"))
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"access_token\":\"%s\",\"refresh_token\":\"%s\"}", access, refresh)))
}
