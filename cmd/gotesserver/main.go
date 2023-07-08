package main

import (
    "github.com/t3mnikov/gotes/internal/app/gotesserver"
    "log"
)

// App start at here
func main() {
    err := gotesserver.Start(gotesserver.NewConfig())
    if err != nil {
        log.Fatal(err)
    }
}
