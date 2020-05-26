package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model.
type User struct {
	gorm.Model
	Username string  `gorm:"unique;unique_index;not null"`
	Password []byte  `gorm:"not null"`
	Noises   []Noise `gorm:"foreignkey:AuthorRefer;association_foreignkey:ID"`
}

// SignUp create users by using username and password.
// Password is hashed by using  bcrypt package.
func SignUp(username string, password string) error {
	err := checkInfo(username, password)
	if err != nil {
		return err
	}

	// Create hashed password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user model.
	user := User{
		Username: username,
		Password: hashedPassword,
	}

	db := DB.Create(&user)
	if err := db.Error; err != nil {
		return fmt.Errorf("Error occur while creating the account: %v", err)
	}

	return nil
}

// Login checks the db with given username and password.
// Returns a result by checking if the user exist.
func Login(username string, password string) error {
	user := User{}
	db := DB.Where("username = ?", username).First(&user)
	if err := db.Error; err != nil {
		return fmt.Errorf("There is no user that named: %v", username)
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return fmt.Errorf("Password is wrong.")
	}

	return nil
}

// DeleteAccount deletes the user.
func DeleteAccount(username string) error {
	user := User{}
	DB.Where("username = ?", username).First(&user)

	db := DB.Delete(&user)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// checkInfo checks if the username and the password is valid.
func checkInfo(username string, password string) error {
	if len(username) < 3 {
		return fmt.Errorf("The username length must be at least 3 characters. \"%v\" is invalid", username)
	}
	if len(password) < 8 {
		return fmt.Errorf("The password length must be at least 8 characters.")
	}

	return nil
}
