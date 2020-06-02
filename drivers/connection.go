package drivers

import (
	"fmt"
	"os"

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

// RedisConnect returns redis connection that is defiend at the env file.
func RedisConnect() (redis.Conn, error) {
	conn, err := redis.Dial("tcp",
		os.Getenv("REDIS_URL"),
		redis.DialClientName(os.Getenv("REDIS_CLIENTNAME")),
		redis.DialPassword(os.Getenv("REDIS_PASSWORD")))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
