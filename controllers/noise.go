package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/boratanrikulu/noisy-notes/models"
)

func NoiseIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user, err := CurrentUser(w, r)
	if err != nil {
		return
	}

	noises, err := models.NoiseIndex(user)
	if err != nil {
		// Return 403. There is an issue with creating noise.
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
		return
	}

	// Return 200. Noises are listed.
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(noises)
}

func NoiseCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user, err := CurrentUser(w, r)
	if err != nil {
		return
	}

	title := r.PostFormValue("title")
	noise, err := models.NoiseCreate(user, title)
	if err != nil {
		// Return 403. There is an issue with creating noise.
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
		return
	}

	// Return 202. Noise is created.
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(noise)
}

func NoiseShow(w http.ResponseWriter, r *http.Request) {
}
