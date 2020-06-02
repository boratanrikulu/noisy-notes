package models

import (
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"

	"github.com/boratanrikulu/noisy-notes/drivers"
)

var (
	name     string
	surname  string
	username string
	password string
	token    string
	title    string
)

// init sets env keys and set db and redis connection..
func init() {
	// Set env keys
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// Set database connection.
	DB, err = drivers.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Set redis connection.
	RPOOL = drivers.RedisPool()
	if err != nil {
		log.Fatal(err)
	}

	// Set the user info to use in testing.
	rand.Seed(time.Now().UnixNano())
	name = randomString(8)
	surname = randomString(8)
	username = randomString(8)
	password = randomString(12)
	title = randomString(12)
}

// randomString returns a random word.
func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
