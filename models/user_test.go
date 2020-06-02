package models

import (
	"testing"
)

// TestSignUp creates an account.
func TestSignUp(t *testing.T) {
	user, err := SignUp(name, surname, username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log("User is created.")
	t.Logf("Username: %v", user.Username)
	t.Logf("Password: %v", string(user.Password))
}

// TestLogin logins and takes a token.
func TestLogin(t *testing.T) {
	resp, err := Login(username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}
	token = resp

	t.Logf("Login is successful for: %v", username)
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
	_, err := SignUp(name, surname, username, password)
	if err == nil {
		t.Fatalf(err.Error())
	}
}

// TestDeleteAccountPermanently deletes testing's user object from DB.
func TestDeletePermanently(t *testing.T) {
	user, err := CurrentUser(token)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = user.DeletePermanently()
	if err != nil {
		t.Fatalf(err.Error())
	}
}
