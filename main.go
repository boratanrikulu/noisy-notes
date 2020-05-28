package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/boratanrikulu/noisy-notes/controllers"
	"github.com/boratanrikulu/noisy-notes/drivers"
	"github.com/boratanrikulu/noisy-notes/models"
)

// init sets env keys, set db and redis connection and makes migrations.
func init() {
	// Set env keys.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Set database connection.
	models.DB, err = drivers.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Set redis connection.
	models.R, err = drivers.RedisConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Make migrations from the schema file.
	db := models.Migrate()
	if err = db.Error; err != nil {
		log.Fatal("Migration is not successful: %v", err)
	}
}

func main() {
	// TODO move route to a separated package.
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Welcome).Methods("GET")
	r.HandleFunc("/recognize", controllers.Recognize).Methods("POST")

	user := r.PathPrefix("/users").Subrouter()
	user.HandleFunc("/login", controllers.Login).Methods("POST")
	user.HandleFunc("/signup", controllers.SignUp).Methods("POST")

	noises := user.PathPrefix("/noises").Subrouter()
	noises.HandleFunc("/", controllers.NoiseIndex).Methods("GET")
	noises.HandleFunc("/", controllers.NoiseCreate).Methods("POST")
	noises.HandleFunc("/{id}", controllers.NoiseShow).Methods("GET")

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
