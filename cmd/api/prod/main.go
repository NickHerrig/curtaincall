package main

import (
    "fmt"
    "net/http"
    "log"

    "curtaincall.tech/pkg/retrieving"

    "curtaincall.tech/pkg/http/rest"
    "curtaincall.tech/pkg/storage/sqlite"
)

func main() {

    var retriever retrieving.Service
    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    retriever = retrieving.NewService(s)
    router := rest.Handler(retriever)

    fmt.Println("The API server is running on port :8888")
    log.Fatal(http.ListenAndServe(":8888", router))
}
