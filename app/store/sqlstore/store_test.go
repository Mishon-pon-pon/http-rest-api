package sqlstore_test

import "testing"

import "os"

var (
	databaseURL string = "host=localhost dbname=restapi_test sslmode=disable password=123"
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable password=123"
	}

	os.Exit(m.Run())
}
