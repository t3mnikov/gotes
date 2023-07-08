package sqlstore

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "strings"
    "testing"
)

// On test DB
func TestDB(t *testing.T, DatabaseURL string) (*sql.DB, func(...string)) {
    t.Helper()

    db, err := sql.Open("postgres", DatabaseURL)
    if err != nil {
        t.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        t.Fatal(err)
    }

    return db, func(tables ...string) {
        if len(tables) > 0 {
            _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
            if err != nil {
                t.Fatal(err)
            }
        }
    }
}
