package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Noise model.
type Noise struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Author      User   `gorm:"foreignkey:AuthorRefer"`
	AuthorRefer uint   `gorm:"not null"`
	Tags        []Tag  `gorm:"foreignkey:NoiseRefer;association_foreignkey:ID"`
}

// SetNoises sets all noises for the user.
func (user *User) SetNoises() error {
	noises := []Noise{}
	db := DB.Order("created_at desc").
		Model(user).
		Association("Noises").
		Find(&noises)

	if err := db.Error; err != nil {
		return err
	}

	user.Noises = noises
	return nil
}

// GetNoise returns wanted noise if it is associated with the user.
func (user *User) GetNoise(id string) (Noise, error) {
	noise := Noise{}
	db := DB.Where("id = ?", id).
		Model(user).
		Association("Noises").
		Find(&noise)

	if err := db.Error; err != nil {
		return Noise{}, err
	}

	return noise, nil
}

// NoiseCreate creates a noise for the user.
func (user *User) NoiseCreate(title string) (Noise, error) {
	noise := Noise{
		Title:  title,
		Author: *user,
	}

	db := DB.Create(&noise)
	if err := db.Error; err != nil {
		return noise, fmt.Errorf("Error occur while creating the noise: %v", err)
	}

	return noise, nil
}

// Delete deletes the noise.
func (noise *Noise) Delete() error {
	db := DB.Delete(noise)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// DeletePermanently deletes the noise permanently.
func (noise *Noise) DeletePermanently() error {
	db := DB.Unscoped().Delete(noise)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}
