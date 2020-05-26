package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/boratanrikulu/noisy-notes/controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.WelcomeGet).Methods("GET")
	r.HandleFunc("/recognize", controllers.RecognizePost).Methods("POST")

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
