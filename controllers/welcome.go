package controllers

import (
	"encoding/json"
	"net/http"
)

type pageResponse struct {
	Error   string
	Message string
}

// WelcomeGet just returns a hello message.
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	res := pageResponse{
		Message: "Hey!",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
