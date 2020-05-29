package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/boratanrikulu/noisy-notes/models"
)

// SignUp creates users
// Form must contains "name", "surname", "username" and "password".
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	name := r.PostFormValue("name")
	surname := r.PostFormValue("surname")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, err := models.SignUp(name, surname, username, password)
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
		User    struct {
			ID        uint
			CreatedAt time.Time
			Name      string
			Surname   string
			Username  string
		}
	}{
		Message: "Account is created.",
		User: struct {
			ID        uint
			CreatedAt time.Time
			Name      string
			Surname   string
			Username  string
		}{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Name:      user.Name,
			Surname:   user.Surname,
			Username:  user.Username,
		},
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
		Token     string
		TokenType string
		ExpiresIn int
	}{
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: 3600,
	})
}

// Delete deletes the account.
func Delete(w http.ResponseWriter, r *http.Request) {
}
