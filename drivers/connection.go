package drivers

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConnect returns a connection between PostgreSQL database that is defined at the env file.
func DBConnect() (*gorm.DB, error) {
	dbinfo := fmt.Sprintf(os.Getenv("DATABASE_URL"))

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// RedisPool returns a pool for Redis.
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 140 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.DialURL(os.Getenv("REDIS_URL")) },
	}
}
