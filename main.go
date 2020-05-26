package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/boratanrikulu/noisy-notes/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.WelcomeGet).Methods("GET")
	r.HandleFunc("/recognize", controllers.RecognizePost).Methods("POST")

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
