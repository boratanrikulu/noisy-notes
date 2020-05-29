package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/boratanrikulu/noisy-notes/models"
)

var (
	CurrentUser models.User
)

// UserAuthMiddleware returns a middleware.
// That middleware check if the user's token is valid.
func UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := getCurrentUser(r)

		if err != nil {
			// Return 401. There is an issue with gettin current user.
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(struct {
				Error string
			}{
				Error: err.Error(),
			})
			return
		}

		// Set current user and call the method.
		CurrentUser = user
		next.ServeHTTP(w, r)
	})
}

// getCurrentUser parses the request to get token.
// Gets the current user by calling models.CurrentUser.
func getCurrentUser(r *http.Request) (models.User, error) {
	headerAuth := r.Header.Get("Authorization")
	authorization := strings.Split(strings.TrimSpace(headerAuth), " ")
	if len(authorization) != 2 {
		return models.User{}, fmt.Errorf("Invalid")
	}

	authType := authorization[0]
	if authType != "Bearer" {
		return models.User{}, fmt.Errorf("Auth token type must be Bearer: %v", authType)
	}

	token := authorization[1]
	user, err := models.CurrentUser(token)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
