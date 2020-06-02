package models

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

// DB and R variables is exported to use on the whole models package.
// The connections is set by using drivers on the main package.
var DB *gorm.DB

// RPOOL variables keeps a pool to get connection when need to access Redis.
var RPOOL *redis.Pool

// Migrate makes migrations by using gorm.
func Migrate() *gorm.DB {
	return DB.AutoMigrate(&User{}, &Noise{}, &NoiseFile{}, &Tag{})
}
