package drivers

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	if got := Connect(); got == nil {
		t.Fatal("Connection is not set..")
	}
}
