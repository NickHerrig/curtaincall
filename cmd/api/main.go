package main

import (
    "fmt"
    "net/http"
    "log"

    "curtaincall.tech/pkg/creating"
    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/deleting"

    "curtaincall.tech/pkg/http/rest"
    "curtaincall.tech/pkg/storage/sqlite"
)

func main() {

    var creator   creating.Service
    var retriever retrieving.Service
    var deleter   deleting.Service
    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    creator = creating.NewService(s)
    retriever = retrieving.NewService(s)
    deleter = deleting.NewService(s)
    router := rest.Handler(creator, retriever, deleter)

    fmt.Println("The API server is running on port :8888")
    log.Fatal(http.ListenAndServe(":8888", router))
}
