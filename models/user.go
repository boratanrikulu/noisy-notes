package models

import (
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
		return err
	}

	return nil
}

// DeleteAccount deletes the user.
func DeleteAccount(username string) error {
	// TODO check the session-cookie.

	user := User{}
	DB.Where("username = ?", username).First(&user)

	db := DB.Delete(&user)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}
