package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tiltedEnmu/puregrade_post/internal/entities"
)

// create is something
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("io.ReadAll() failed: /create", err.Error())
		_, _ = w.Write([]byte("Reading error"))
		return
	}

	var input entities.Post
	if err := json.Unmarshal(buf, &input); err != nil {
		log.Println("json.Unmarshal() failed: /create ", err.Error())
		_, _ = w.Write([]byte("Unmarshal failed!"))
		return
	}

	id, err := h.services.CreatePost(&input)
	if err != nil {
		log.Println("CreatePost() failed: /create ", err.Error())
		_, _ = w.Write([]byte("Post creation failed"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf(`{"id":"%s"}`, id)))
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	post, err := h.services.GetPost(id)
	if err != nil {
		log.Println("Getting post failed: /get ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not-found"))
		return
	}

	marshaled, err := json.Marshal(post)
	if err != nil {
		log.Println("Marshal failed: /get ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf(`{"post":"%s"}`, marshaled)))
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := h.services.DeletePost(id)
	if err != nil {
		log.Println("Deleting post failed: /delete ", err.Error())
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Deleting post failed"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf(`{"id":"%s"}`, id)))
}
