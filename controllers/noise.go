package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// NoiseIndex returns all noise for the current user.
func NoiseIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	noises, err := CurrentUser.GetNoises()
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

// NoiseCreate create a noise for the current user.
// Form must contains "title" and "file".
// Form may contains "tags". Example: "Tag 1, Tag 2"
func NoiseCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := getNoiseFile(r)
	if err != nil {
		// Return 403. There is an issue with getting noise file.
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
		return
	}

	title := r.PostFormValue("title")
	tags := getTagsFromString(r.PostFormValue("tags"))

	noise, err := CurrentUser.NoiseCreate(title, b.Bytes(), tags)
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

// NoiseShow returns the noise.
// There must be an {id} parameter.
func NoiseShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	noise, err := CurrentUser.GetNoise(id)
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

	// Return 200. Noise is listed.
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(noise)
}

// NoiseFileShow returns the noise's file.
// There must be an {id} parameter.
func NoiseFileShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	noise, err := CurrentUser.GetNoiseWithFile(id)
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

	resp := bytes.NewReader(noise.File.Data)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v.mp3\"", noise.Title))
	w.Header().Set("Content-Type", "audio/mpeg;")

	io.Copy(w, resp)
}

// getNoiseFile returns a buffer for the file from the request.
func getNoiseFile(r *http.Request) (*bytes.Buffer, error) {
	// Set file
	r.ParseMultipartForm(1 << 20) // 10 MB size limit
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b bytes.Buffer
	io.Copy(&b, file)

	// Check the format.
	err = checkMime(b.Bytes())
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// getTagsFromString returns a string array
// by checking, clearing and uniqunig the given string.
func getTagsFromString(s string) []string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil
	}

	tags := []string{}
	for _, tag := range strings.Split(s, ",") {
		tag = strings.TrimSpace(tag)
		if (len(tag) == 0) || contains(tags, tag) {
			continue
		}
		tags = append(tags, tag)
	}

	return tags
}

// checkMime checks if the file's mime is allowed.
func checkMime(b []byte) error {
	allowed := []string{"audio/mpeg"}
	mime := http.DetectContentType(b)
	if contains(allowed, mime) {
		return nil
	}
	return fmt.Errorf("Not allowed format: %v\nPlease upload one of the this formats: %v",
		mime, strings.Join(allowed, ", "))
}

// contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
