package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Welcome to Curtain Call, your Digital Theater Experience!"))
}


func GetShow(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to <SHOW>!"))
}

func main () {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/show", GetShow)

    log.Println("Starting the application on port :8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
