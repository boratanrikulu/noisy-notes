package models

import (
	"fmt"
	"log"

	"github.com/boratanrikulu/noisy-notes/noises"
	"github.com/jinzhu/gorm"
)

// Noise model.
type Noise struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Author      User      `gorm:"foreignkey:AuthorRefer" json:"-"`
	AuthorRefer uint      `gorm:"not null" json:"-"`
	Tags        []Tag     `gorm:"many2many:noise_tags"`
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

// GetNoises returns all noises that are matched by checking
// noises' titles, noises' texts and tags' titles.
// It can set sort type and limit.
func (user *User) GetNoises(q string, sort string, take int) ([]Noise, error) {
	noises := []Noise{}

	db := DB.Order("updated_at "+sort).
		Limit(take).
		Preload("Tags").
		Where("noises.title ILIKE ? OR noises.text ILIKE ? OR EXISTS(?)",
			"%%"+q+"%%",
			"%%"+q+"%%",
			DB.Table("tags").
				Joins("JOIN noise_tags ON noise_tags.noise_id = noises.id").
				Where("tags.id = noise_tags.tag_id AND tags.title ILIKE ?",
					"%%"+q+"%%").
				SubQuery()).
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
		Preload("Tags").
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
		Preload("File").
		Model(user).
		Association("Noises").
		Find(&noise)
	if err := db.Error; err != nil {
		return Noise{}, err
	}

	return noise, nil
}

// NoiseCreate creates a noise for the user.
func (user *User) NoiseCreate(title string, file []byte, ts []string) (Noise, error) {
	noise := Noise{
		Title: title,
		File: NoiseFile{
			Data: file,
		},
	}

	// set tags.
	tags, err := getFirstOrCreateTags(user, ts)
	if err != nil {
		return Noise{}, err
	}
	if len(tags) > 0 {
		noise.Tags = tags
	}

	// create user's noise.
	db := DB.Model(user).Association("Noises").Append(&noise)
	if err := db.Error; err != nil {
		return Noise{}, fmt.Errorf("Error occur while creating the noise: %v", err)
	}

	return noise, nil
}

// AfterCreate runs starting to recognition and returns.
// Recognition ops are run as a goroutine.
func (noise *Noise) AfterCreate(scope *gorm.Scope) error {
	// result channel
	c := make(chan string)
	// error channel
	e := make(chan error)

	// start to recognition
	go noises.Recognize(noise.File.Data, c, e)
	// start  to setting results
	go setResults(DB, noise, c, e)

	// return, just return!
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

// setResults set the result that are received from the channel.
func setResults(DB *gorm.DB, noise *Noise, c chan string, e chan error) {
	select {
	case err := <-e:
		log.Println(err.Error())
		close(c)
		close(e)
	case text := <-c:
		noise.IsActive = true
		noise.Text = text
		db := DB.Save(noise)
		if err := db.Error; err != nil {
			log.Println(err.Error())
		}
		close(c)
		close(e)
	}
}

// Delete deletes the noise.
func (noise *Noise) Delete() error {
	db := DB.Delete(noise)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// getFirstorCreateTags returns tag array that is created
// by checing if the user has the tag or not.
func getFirstOrCreateTags(user *User, ts []string) ([]Tag, error) {
	tags := []Tag{}

	for _, t := range ts {
		tag := Tag{}

		db := DB.Where("title = ?", t).
			Model(user).
			Association("Tags").
			Find(&tag)

		// create the tag if user has not.
		if err := db.Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				tag.Title = t
				db := DB.Model(user).Association("Tags").Append(&tag)
				if err := db.Error; err != nil {
					return nil, fmt.Errorf("Error occur while creating the noise: %v", err)
				}
			} else {
				return nil, err
			}
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
