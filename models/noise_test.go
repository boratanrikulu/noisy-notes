package models

import (
	"testing"
)

// TestNoiseCreate creates a account, login with it.
// Creates a noise for the user.
// Removes the noise, removes the user.
func TestNoiseCreate(t *testing.T) {
	user, err := SignUp(name, surname, username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log("User is created.")
	t.Logf("Username: %v", user.Username)
	t.Logf("Password: %v", string(user.Password))

	token, err := Login(username, password)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Login is successful: %v", token)

	user, err = CurrentUser(token)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Current user: %v", user.Username)

	noise, err := user.NoiseCreate(title)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Noise is created: %v", noise.Title)

	err = noise.DeletePermanently()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("The noise is deleted.")

	err = user.DeletePermanently()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("The account is deleted.")

}
