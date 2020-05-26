package controllers

import (
	"fmt"
	"net/http"

	"github.com/boratanrikulu/noisy-notes/models"
)

// SignUp creates users
//
// Form must contains "username" and "password".
func SignUp(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	err := models.SignUp(username, password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Account is created.")
}

// Login creates session and returns a token.
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	err := models.Login(username, password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Login is successful.")
}

// Delete deletes the account.
func Delete(w http.ResponseWriter, r *http.Request) {
}
