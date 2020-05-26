package drivers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect returns a connection between PostgreSQL database that is defined at the env file.
func Connect() *gorm.DB {
	dbinfo := fmt.Sprintf("host=%v port=%v user=%v password=%v, dbname=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err.Error())
	}

	return db
}
