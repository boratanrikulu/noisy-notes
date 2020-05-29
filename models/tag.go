package models

import (
	"github.com/jinzhu/gorm"
)

// Tag model.
type Tag struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Author      User   `gorm:"foreignkey:AuthorRefer"`
	AuthorRefer uint   `gorm:"not null"`
	Noise       Noise  `gorm:"foreignkey:NoiseRefer"`
	NoiseRefer  uint   `gorm:"not null"`
}
