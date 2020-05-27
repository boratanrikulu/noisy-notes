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
	token    string
)

// init sets env keys and set db and redis connection..
func init() {
	// Set env keys
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	DB, err = drivers.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Set redis connection.
	R, err = drivers.RedisConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Set the user info to use in testing.
	rand.Seed(time.Now().UnixNano())
	username = randomString(8)
	password = randomString(12)
}

// TestSignUp creates an account.
func TestSignUp(t *testing.T) {
	err := SignUp(username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log("User is created.")
	t.Logf("Username: %v", username)
	t.Logf("Password: %v", password)
}

// TestLogin logins and takes a token.
func TestLogin(t *testing.T) {
	resp, err := Login(username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}
	token = resp

	t.Logf("Login is succesful for: %v", username)
	t.Logf("Token: %v", token)
}

// TestCurrentUser takes the current user with the token.
func TestCurrentUser(t *testing.T) {
	currentUser, err := CurrentUser(token)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Logf("Current user is %v", currentUser.Username)
}

// TestCurrentUserWithWrongToken tries to access user with a wrong token.
func TestCurrentUserWithWrongToken(t *testing.T) {
	_, err := CurrentUser("randomtoken-invalid-token")
	if err == nil {
		t.Fatalf(err.Error())
	}
}

// TestDuplicatedUsernames creates two account with the same username.
// It expect an error from the database.
func TestDuplicatedUsernames(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	username := randomString(8)
	password := randomString(12)

	err := SignUp(username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = SignUp(username, password)
	if err == nil {
		t.Fatalf(err.Error())
	}
}

// TestDeleteAccount checks the DeleteAccount method,
// by deleting the user that is created on testing.
func TestDeleteAccount(t *testing.T) {
	err := DeleteAccount(username)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%v is deleted.", username)
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
