package main

import (
    "log"
    "net/http"
    "time"

    "curtaincall.tech/pkg/web"
    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/storage/sqlite"

    "github.com/gorilla/mux"
    "github.com/justinas/alice"


)

func main() {

    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    r := retrieving.NewService(s)

    standardMiddleware := alice.New(web.RecoverPanic, web.SecureHeaders, web.CorsHeaders)    

    rtr := mux.NewRouter()
    api := rtr.PathPrefix("/api/v1/").Subrouter()
    api.HandleFunc("/shows", web.RetrieveAllShows(r)).Methods("GET")


	handler := standardMiddleware.Then(rtr)

	srv := &http.Server{
		Addr:         ":8888",
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Listening on :8888")
	log.Fatal(srv.ListenAndServe())

}
