package models

import (
	"github.com/jinzhu/gorm"
)

// Migrate makes migrations by using gorm.
func Migrate(db *gorm.DB) *gorm.DB {
	return db.AutoMigrate(&User{}, &Noise{})
}
