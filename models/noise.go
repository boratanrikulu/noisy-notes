package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/boratanrikulu/noisy-notes/noises"
)

// Noise model.
type Noise struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Author      User      `gorm:"foreignkey:AuthorRefer" json:"-"`
	AuthorRefer uint      `gorm:"not null" json:"-"`
	Tags        []Tag     `gorm:"foreignkey:NoiseRefer;association_foreignkey:ID"`
	File        NoiseFile `gorm:"foreignkey:NoiseRefer;association_foreignkey:ID" json:"-"`
	Text        string
	IsActive    bool `gorm:"default:false"`
}

// NoiseFile model.
type NoiseFile struct {
	gorm.Model
	Data       []byte `gorm:"not null"` // TODO move file to gcloud.
	Noise      *Noise `gorm:"foreignkey:NoiseRefer"`
	NoiseRefer uint   `gorm:"not null"`
}

// GetNoises returns all noises for the user.
func (user *User) GetNoises() ([]Noise, error) {
	noises := []Noise{}
	db := DB.Order("created_at desc").
		Model(user).
		Association("Noises").
		Find(&noises)

	if err := db.Error; err != nil {
		return nil, err
	}

	return noises, nil
}

// GetNoise returns wanted noise that is associated with the user.
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

// GetNoiseWithFile returns wanted noise's file that is associated with the user.
func (user *User) GetNoiseWithFile(id string) (Noise, error) {
	noise := Noise{}

	db := DB.Where("id = ?", id).
		Model(user).
		Association("Noises").
		Find(&noise)
	if err := db.Error; err != nil {
		return Noise{}, err
	}

	db = DB.Model(&noise).Association("File").Find(&noise.File)
	if err := db.Error; err != nil {
		return Noise{}, err
	}

	return noise, nil
}

// NoiseCreate creates a noise for the user.
func (user *User) NoiseCreate(title string, file []byte) (Noise, error) {
	noise := Noise{
		Title: title,
		File: NoiseFile{
			Data: file,
		},
	}

	db := DB.Model(user).Association("Noises").Append(&noise)
	if err := db.Error; err != nil {
		return Noise{}, fmt.Errorf("Error occur while creating the noise: %v", err)
	}

	return noise, nil
}

// AfterCreate runs Recognize method after the creation.
// Set the returned Text value to the Noise.
func (noise *Noise) AfterCreate(scope *gorm.Scope) error {
	text, err := noises.Recognize(noise.File.Data)
	if err != nil {
		return err
	}

	noise.IsActive = true
	noise.Text = text
	db := scope.DB().Save(noise)
	if err := db.Error; err != nil {
		return err
	}

	return nil
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
