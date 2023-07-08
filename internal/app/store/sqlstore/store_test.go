package sqlstore_test

import (
    "os"
    "testing"
)

var (
    DatabaseURL string
)

func TestMain(m *testing.M) {
    DatabaseURL = os.Getenv("DATABASE_URL")
    if DatabaseURL == "" {
        DatabaseURL = "host=localhost user=user password=123456789 dbname=gotes_test sslmode=disable"
    }

    os.Exit(m.Run())
}
