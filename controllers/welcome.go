package controllers

import (
	"fmt"
	"net/http"
)

// WelcomeGet just returns a hello message.
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := "{ \"response\":\"Hello World.\" }"
	fmt.Fprint(w, response)
}
