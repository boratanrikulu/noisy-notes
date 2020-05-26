package models

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/boratanrikulu/noisy-notes/drivers"
)

var (
	username string
	password string
)

func init() {
	// Set env keys
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	DB = drivers.Connect()

	// Set the user info to use in testing.
	rand.Seed(time.Now().UnixNano())
	username = randomString(8)
	password = randomString(12)
}

// TestSignUp creates an account to test SignUp method.
func TestSignUp(t *testing.T) {
	err := SignUp(username, password)
	if err != nil {
		t.Fatalf("Error occur while creating user: %v", err)
	}

	t.Log("User is created.")
	t.Logf("Username: %v", username)
	t.Logf("Password: %v", password)
}

// TestDuplicatedUsernames creates two account with the same username.
// It expect an error from the database.
func TestDuplicatedUsernames(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	username := randomString(8)
	password := randomString(12)

	err := SignUp(username, password)
	if err != nil {
		t.Fatalf("Error occur while creating user: %v", err)
	}

	err = SignUp(username, password)
	if err == nil {
		t.Fatalf("Error occur. We created two user account with the same username")
	}
}

// TestDeleteAccount checks the DeleteAccount method,
// by deleting the user that is created on testing.
func TestDeleteAccount(t *testing.T) {
	err := DeleteAccount(username)
	if err != nil {
		t.Fatal("Error occur while deleting the user.")
	}
	t.Logf("%v is deleted.", username)
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
