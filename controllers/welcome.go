package controllers

import (
	"encoding/json"
	"net/http"
)

// WelcomeGet just returns a hello message.
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct {
		Message string
	}{
		Message: "Hey!",
	})
}
