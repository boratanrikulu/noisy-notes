package drivers

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func TestConnect(t *testing.T) {
	if got := Connect(); got == nil {
		t.Fatal("Connection is not set..")
	}
}
