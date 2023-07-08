package main

import (
    "database/sql"
    "flag"
    _ "github.com/lib/pq"
    "github.com/pressly/goose/v3"
    "github.com/t3mnikov/gotes/internal/app/gotesserver"
    "log"
    "strings"
)

// Migration scripted application
func main() {
    command := flag.String("c", "status", "command")
    dir := flag.String("dir", "migrations", "migrations directory")
    args := flag.String("args", "", "command arguments")
    flag.Parse()

    config := gotesserver.NewConfig()
    dsn := config.DatabaseURL
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("-dbstring=%q: %v\n", dsn, err)
    }

    defer db.Close()

    err = goose.SetDialect("postgres")
    if err != nil {
        log.Fatal(err)
    }

    argsValues := strings.Split(*args, " ")

    err = goose.Run(*command, db, *dir, argsValues...)
    if err != nil {
        log.Fatalf("goose run: %v", err)
    }
}
