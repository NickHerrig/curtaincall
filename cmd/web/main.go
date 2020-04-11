package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))


    log.Println("Starting Curtain Call on port :9000")
    err := http.ListenAndServe(":9000", mux)
    log.Fatal(err)
}
