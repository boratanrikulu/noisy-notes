package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"

	"github.com/boratanrikulu/noisy-notes/controllers/helpers"
	"github.com/boratanrikulu/noisy-notes/noises"
)

// NoiseIndex returns all noise for the current user.
func NoiseIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	q, sort, take, err := getParamsFromRequest(r)
	if err != nil {
		// Return 403. There is an issue with creating noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	noises, err := CurrentUser.GetNoises(q, sort, take)
	if err != nil {
		// Return 403. There is an issue with creating noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
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
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	title := r.PostFormValue("title")
	tags := getTagsFromString(r.PostFormValue("tags"))

	noise, err := CurrentUser.NoiseCreate(title, b, tags)
	if err != nil {
		// Return 403. There is an issue with creating noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
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
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
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
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	resp := bytes.NewReader(noise.File.Data)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v.wav\"", noise.Title))
	w.Header().Set("Content-Type", "audio/wav;")

	io.Copy(w, resp)
}

// NoiseUpdate update the noise.
// It works just like creating.
func NoiseUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	noise, err := CurrentUser.GetNoise(id)
	if err != nil {
		// Return 403. There is an issue with getting the noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	b, err := getNoiseFile(r)
	if err != nil {
		// Return 403. There is an issue with getting noise file.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	title := r.PostFormValue("title")
	tags := getTagsFromString(r.PostFormValue("tags"))

	err = noise.Update(&CurrentUser, title, b, tags)
	if err != nil {
		// Return 403. There is an issue with updating noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	// Return 202. Noise is update.
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(noise)
}

// NoiseDelete temporarily deletes the given noise and it's file
func NoiseDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	noise, err := CurrentUser.GetNoise(id)
	if err != nil {
		// Return 403. There is an issue with getting noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	err = noise.Delete()
	if err != nil {
		// Return 403. There is an issue with creating noise.
		helpers.ReturnError(w, http.StatusForbidden, err.Error())
		return
	}

	// Return 200. Noise is deleted.
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct {
		Message string
	}{
		Message: "The noise is deleted.",
	})
}

// getNoiseFile returns a buffer for the file from the request.
func getNoiseFile(r *http.Request) ([]byte, error) {
	// read the file.
	r.ParseMultipartForm(1 << 20) // 10 MB size limit
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// get the mime from file.
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return nil, fmt.Errorf("We could not find file's mime.")
	}
	file.Seek(0, 0)

	// check the mime if is allowed.
	allowed := []string{"audio/mpeg",
		"audio/mp3",
		"audio/ogg",
		"audio/wav",
		"audio/flac",
		"audio/aac"}
	if !mimetype.EqualsAny(mime.String(), allowed...) {
		return nil, fmt.Errorf("Not allowed format: %v\nPlease upload one of the this formats: %v",
			mime.String(), strings.Join(allowed, ", "))
	}

	// convert file to wav.
	cB, err := noises.Convert(file)
	if err != nil {
		return nil, err
	}

	return cB, nil
}

func getParamsFromRequest(r *http.Request) (q string, sort string, take int, err error) {
	// query, it may be a word or a sentence..
	q = strings.TrimSpace(r.FormValue("q"))

	// sort, only allowed "asc" and "desc"
	sort = strings.ToLower(strings.TrimSpace(r.FormValue("sort")))
	if sort == "" {
		sort = "desc"
	}
	if sort != "asc" && sort != "desc" {
		err = fmt.Errorf("Sort must be \"asc\" or \"desc\"")
		return
	}

	// take, must be an integer
	t := r.FormValue("take")
	take = -1 // set default if take is not exist.
	if t != "" {
		take, err = strconv.Atoi(t)
	}
	if err != nil {
		err = fmt.Errorf("Take must be a integer")
		return
	}

	return
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

// contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
