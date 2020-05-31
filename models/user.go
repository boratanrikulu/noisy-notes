package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model.
type User struct {
	gorm.Model
	Name     string  `gorm:"not null"`
	Surname  string  `gorm:"not null"`
	Username string  `gorm:"unique;unique_index;not null"`
	Password []byte  `gorm:"not null" json:"-"`
	Noises   []Noise `gorm:"foreignkey:AuthorRefer;association_foreignkey:ID"`
	Tags     []Tag   `gorm:"foreignkey:AuthorRefer;association_foreignkey:ID"`
}

// SignUp creates users by using username and password.
// Password is hashed by using  bcrypt package.
func SignUp(name string, surname string, username string, password string) (User, error) {
	err := checkInfo(name, surname, username, password)
	if err != nil {
		return User{}, err
	}

	// Create hashed password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// Create user model.
	user := User{
		Name:     name,
		Surname:  surname,
		Username: username,
		Password: hashedPassword,
	}

	db := DB.Create(&user)
	if err := db.Error; err != nil {
		return User{}, fmt.Errorf("Error occur while creating the account: %v", err)
	}

	return user, nil
}

// Login checks the db with given username and password.
// Returns a result by checking if the user exist.
func Login(username string, password string) (string, error) {
	user := User{}

	// Get the user from the DB.
	db := DB.Where("username = ?", username).First(&user)
	if err := db.Error; err != nil {
		return "", fmt.Errorf("The user is not exist: %v", username)
	}

	// Check the if password is correct.
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return "", fmt.Errorf("Password is not correct.")
	}

	// Create a session that lives for 1 hour.
	// key: token, value: username.
	token := uuid.New().String()
	_, err = R.Do("SETEX", token, 3600, username)
	if err != nil {
		return "", fmt.Errorf("Error occur while creating session: %v", err)
	}

	// Return the token.
	return token, nil
}

// Logout removes the current user's session.
func (user *User) Logout(token string) error {
	_, err := R.Do("DEL", token)
	if err != nil {
		return fmt.Errorf("Error occur while deleting the session.")
	}

	// Clear the current user.
	user = nil

	return nil
}

// CurrentUser returns the current user that matches with the token.
func CurrentUser(token string) (User, error) {
	user := User{}

	// Get the username if redis has the token.
	resp, err := R.Do("GET", token)
	if err != nil || resp == nil || resp == "" {
		return User{}, fmt.Errorf("Authorization is not permitted.")
	}

	// There is an username. Take the user object from the DB.
	db := DB.Where("username = ?", resp).First(&user)
	if err := db.Error; err != nil {
		return User{}, fmt.Errorf("The user is not exist: %v", err)
	}

	return user, nil
}

// Delete deletes the user.
func (user *User) Delete() error {
	db := DB.Delete(user)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// DeletePermanently deletes the user.
func (user *User) DeletePermanently() error {
	db := DB.Unscoped().Delete(user)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// checkInfo checks if the username and the password is valid.
func checkInfo(name string, surname string, username string, password string) error {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(surname) == "" {
		return fmt.Errorf("The name or surname can not be empty")
	}
	if len(username) < 3 {
		return fmt.Errorf("The username length must be at least 3 characters. \"%v\" is invalid", username)
	}
	if len(password) < 8 {
		return fmt.Errorf("The password length must be at least 8 characters.")
	}

	return nil
}
