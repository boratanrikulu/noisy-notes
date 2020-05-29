package controllers

import (
	"encoding/json"
	"net/http"
)

func NoiseIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := CurrentUser.SetNoises()
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
	_ = json.NewEncoder(w).Encode(CurrentUser)
}

func NoiseCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	title := r.PostFormValue("title")
	noise, err := CurrentUser.NoiseCreate(title)
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
