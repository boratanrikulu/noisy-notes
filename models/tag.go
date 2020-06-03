package models

import (
	"github.com/jinzhu/gorm"
)

// Tag model.
type Tag struct {
	gorm.Model
	Title       string  `gorm:"not null"`
	Author      User    `gorm:"foreignkey:AuthorRefer" json:"-"`
	AuthorRefer uint    `gorm:"not null" json:"-"`
	Noises      []Noise `gorm:"many2many:noise_tags" json:"-"`
}
