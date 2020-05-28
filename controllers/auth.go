package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/boratanrikulu/noisy-notes/models"
)

// CurrentUser calls the getCurrentUser to get current user.
// Returns 401 if there is no current user.
func CurrentUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	user, err := getCurrentUser(r)

	if err != nil {
		// Return 401. There is an issue with gettin current user.
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
		return user, err
	}

	return user, nil
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
