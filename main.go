package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/boratanrikulu/noisy-notes/controllers"
	"github.com/boratanrikulu/noisy-notes/drivers"
	"github.com/boratanrikulu/noisy-notes/models"
)

// DB variable is exported to use on the whole project.
// Connection is set by using driver/connection.
var DB *gorm.DB

func main() {
	// Set env keys.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Set database connection.
	DB = drivers.Connect()

	// Make migrations from the schema file.
	DB = models.Migrate(DB)

	// TODO move route to a separated package.
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.WelcomeGet).Methods("GET")
	r.HandleFunc("/recognize", controllers.RecognizePost).Methods("POST")

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
