package helpers

import (
	"encoding/json"
	"net/http"
)

// ReturnError writes given error message and http status to ResponseWriter.
func ReturnError(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(struct {
		Error string
	}{
		Error: err,
	})
}
