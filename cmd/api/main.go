package main

import (
    "fmt"
    "net/http"
    "log"

    "curtaincall.tech/pkg/adding"
    "curtaincall.tech/pkg/http/rest"
    "curtaincall.tech/pkg/storage/sqlite"

)

func main() {

    var adder adding.Service
    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    adder = adding.NewService(s)
    router := rest.Handler(adder)

    fmt.Println("The API server is running on port :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
