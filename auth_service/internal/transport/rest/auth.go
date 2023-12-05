package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type signInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// This func is a wrapper over the method for creating a new user profile
// and also creates refresh and access tokens and returns them.
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Making a request to UserService and decoding the response if successful
	client := &http.Client{}
	resp, err := client.Post(
		h.UserServiceAddr,
		"application/x-www-form-urlencoded",
		strings.NewReader(string(bytes)),
	)
	if err != nil {
		log.Print("client.Get() failed: /sign-up ", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("User registration failed"))
		return
	}
	if resp.StatusCode != 200 {
		log.Print("client.Get() returned the status code of a failed request: /sign-up ", resp.StatusCode)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("User registration failed"))
		return
	}
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Print("client.Get() returned invalid response: /sign-up", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal bytes to userID
	userID := struct {
		Value string `json:"userID"`
	}{}
	err = json.Unmarshal(bytes, &userID)
	if err != nil {
		log.Print("client.Get() returned invalid response: /sign-up", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate tokens and send back if successful
	access, refresh, err := h.services.GenerateTokens(userID.Value)
	if err != nil {
		log.Print("GenerateTokens() failed: /sign-up ", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Generate tokens failed"))
		return
	}

	_, _ = w.Write(
		[]byte(fmt.Sprintf(
			`{"access_token":"%s","refresh_token":"%s"}`, access, refresh,
		)),
	)
}

// This func receives Email and Password and return access and refresh tokens or error
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If the data is incorrect, we immediately return an error without disturbing the UserService
	var input signInDTO
	if err := json.Unmarshal(bytes, &input); err != nil {
		log.Print("json.Unmarshal() failed: /sign-in ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Making a request to UserService and decoding the response if successful
	req, err := http.NewRequest(
		http.MethodGet,
		h.UserServiceAddr,
		strings.NewReader(string(bytes)),
	)
	if err != nil {
		log.Print("http.NewRequest() failed: /sign-in ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("client.Get() failed: /sign-in ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != 200 {
		log.Print("client.Get() returned the status code of a failed request: /sign-in ", resp.StatusCode)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Print("client.Get() returned invalid response: /sign-in", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate tokens and send back if successful
	access, refresh, err := h.services.GenerateTokens(string(bytes))
	if err != nil {
		log.Print("GenerateTokens() failed: /sign-in ", err)
		w.WriteHeader(500)
		return
	}

	_, _ = w.Write(
		[]byte(fmt.Sprintf(
			`{"access_token":"%s","refresh_token":"%s"}`, access, refresh,
		)),
	)
}

// This func receives refresh token and return access and refresh tokens or error
func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := struct {
		Value string `json:"refresh"`
	}{}
	if err := json.Unmarshal(bytes, &input); err != nil {
		log.Print("json.Unmarshal() failed: /sign-in ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Search for a Refresh token and delete it when found
	id, err := h.services.GetUserId(input.Value)
	if err != nil {
		log.Print("Searching token failed: /refresh ", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("Refresh token was not found"))
		return
	}

	access, refresh, err := h.services.GenerateTokens(id)
	if err != nil {
		log.Print("Generate tokens failed: /refresh ", err.Error())
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Generate tokens failed"))
		return
	}

	_, _ = w.Write(
		[]byte(fmt.Sprintf(
			`{"access_token":"%s","refresh_token":"%s"}`, access, refresh,
		)),
	)
}
