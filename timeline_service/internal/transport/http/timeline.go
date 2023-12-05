package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) getLatest(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	postId := r.URL.Query().Get("postId")

	posts, err := h.services.GetLatestById(userId, postId, 0)
	if err != nil {
		log.Println("Getting timeline failed: /get-latest ", err.Error())
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Getting timeline failed"))
		return
	}

	marshaled, err := json.Marshal(posts)
	if err != nil {
		log.Println("Marshal failed: /get-latest ", err.Error())
		_, _ = w.Write([]byte("Marshal failed"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf(`{"timeline":"%s"}`, marshaled)))
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")

	posts, err := h.services.GetRange(userId, 0, -1)
	if err != nil {
		log.Println("Getting timeline failed: /get ", err.Error())
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Getting timeline failed"))
		return
	}

	marshaled, err := json.Marshal(posts)
	if err != nil {
		log.Println("Marshal failed: /get ", err.Error())
		_, _ = w.Write([]byte("Marshal failed"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf(`{"timeline":"%s"}`, marshaled)))
}
