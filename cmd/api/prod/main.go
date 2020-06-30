package main

import (
    "log"
    "net/http"
    "time"

    "curtaincall.tech/pkg/web"
    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/storage/sqlite"

    "github.com/bmizerany/pat"
    "github.com/justinas/alice"


)

func main() {

    //TODO: Add go environment varialbe importing for api port and fqdn

    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    r := retrieving.NewService(s)

    standardMiddleware := alice.New(web.RecoverPanic, web.SecureHeaders, web.CorsHeaders)    

    //TODO: Add serving of index.html for VUE JS and vue router routes
    //TODO: Look into subroutes for entry point 

    m := pat.New()
    m.Get("/shows", http.HandlerFunc(web.RetrieveAllShows(r)))

	handler := standardMiddleware.Then(m)

    //TODO: Add listener ports from socket file

    //TODO: Add HTTP to HTTPS redirect

    //TODO: Add HTTPS Server details

    //TODO: Add auto cert renewal 

    //TODO: Zero Downtime Deployment?

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
