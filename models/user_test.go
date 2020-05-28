package models

import (
	"testing"
)

// TestSignUp creates an account.
func TestSignUp(t *testing.T) {
	err := SignUp(name, surname, username, password)
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
	err := SignUp(name, surname, username, password)
	if err == nil {
		t.Fatalf(err.Error())
	}
}

// TestDeleteAccountPermanently deletes testing's user object from DB.
func TestDeleteAccountPermanently(t *testing.T) {
	user, err := CurrentUser(token)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = DeleteAccountPermanently(user)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
