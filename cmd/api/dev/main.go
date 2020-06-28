package main

import (
    "fmt"
    "net/http"
    "log"

    "curtaincall.tech/pkg/web"
    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/storage/sqlite"

    "github.com/bmizerany/pat"

)

func main() {

    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    r := retrieving.NewService(s)

    m := pat.New()
    m.Get("/shows", http.HandlerFunc(web.RetrieveAllShows(r)))

    fmt.Println("The API server is running on port :8888")
    log.Fatal(http.ListenAndServe(":8888", m))
}
