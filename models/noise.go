package models

import (
	"github.com/jinzhu/gorm"
)

// Noise model.
type Noise struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Author      User   `gorm:"foreignkey:AuthorRefer"`
	AuthorRefer uint   `gorm:"not null"`
}
