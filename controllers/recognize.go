package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/boratanrikulu/noisy-notes/noises"
)

// RecognizePost gets a file from the response.
// It gets text from the speech.
// It returns the text as a json response.
//
// The file must be named as "noise".
// The size limit of the file is 10 mb.
// The file must be an "audio/mpeg".
func RecognizePost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20) // 10 MB size limit
	noise, _, err := r.FormFile("noise")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}
	defer noise.Close()

	var b bytes.Buffer
	io.Copy(&b, noise)

	// Check the format.
	err = checkMime(b.Bytes())
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	resp, err := noises.Recognize(b.Bytes())
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, resp)
}

// Private methods

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
