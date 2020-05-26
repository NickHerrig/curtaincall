package main

import (
    "fmt"
    "net/http"
    "log"

    "curtaincall.tech/pkg/creating"
    "curtaincall.tech/pkg/http/rest"
    "curtaincall.tech/pkg/storage/sqlite"

)

func main() {

    var creator creating.Service
    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    creator = creating.NewService(s)
    router := rest.Handler(creator)

    fmt.Println("The API server is running on port :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
