package models

import (
	"github.com/jinzhu/gorm"
)

// DB variable is exported to use on the whole package.
// Connection is set by using drivers.Connect()
var DB *gorm.DB

// Migrate makes migrations by using gorm.
func Migrate() *gorm.DB {
	return DB.AutoMigrate(&User{}, &Noise{})
}
