package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=christina password=90909090 host=localhost port=5432 dbname=restapi_test sslmode=disable"
	}
	os.Exit(m.Run())
}
