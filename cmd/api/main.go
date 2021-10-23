package main

import (
	"net/http"
	"log"
	"os"
	"time"
)

type application struct {
	env string
}

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		env: "Development",
	}

	srv := &http.Server{
		Addr: ":3000",
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting server on port :3000")

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
