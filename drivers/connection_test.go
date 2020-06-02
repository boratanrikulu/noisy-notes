package drivers

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// init set env keys.
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Env load failed: %v", err)
	}
}

func TestDBConnect(t *testing.T) {
	if _, err := DBConnect(); err != nil {
		t.Fatalf("DB connection is failed: %v", err)
	}
}

func TestRedisPool(t *testing.T) {
	p := RedisPool()
	if _, err := p.Dial(); err != nil {
		t.Fatalf("Redis connection is failed: %v", err)
	}
}
