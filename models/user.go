package models

import (
	"github.com/jinzhu/gorm"
)

// User model.
type User struct {
	gorm.Model
	Username       string  `gorm:"unique;unique_index;not null"`
	HashedPassword string  `gorm:"not null"`
	Noises         []Noise `gorm:"foreignkey:AuthorRefer;association_foreignkey:ID"`
}
