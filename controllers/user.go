package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/boratanrikulu/noisy-notes/models"
)

// SignUp creates users
// Form must contains "username" and "password".
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	err := models.SignUp(username, password)
	if err != nil {
		// Return 403. There is an issue with creating account.
		w.WriteHeader(http.StatusForbidden)

		_ = json.NewEncoder(w).Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
		return
	}

	// Return 202. Account is created.
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(struct {
		Message string
	}{
		Message: "Account is created.",
	})
}

// Login creates session and returns a token.
// Form must contains "username" and "password".
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	token, err := models.Login(username, password)
	if err != nil {
		// Return 403. There is an issue with login.
		w.WriteHeader(http.StatusForbidden)

		_ = json.NewEncoder(w).Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
		return
	}

	// Return 202. Login is successful.
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(struct {
		Message string
		Token   string
	}{
		Message: "Login is successful.",
		Token:   token,
	})
}

// Delete deletes the account.
func Delete(w http.ResponseWriter, r *http.Request) {
}
