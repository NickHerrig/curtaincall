package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "path"
    "time"

    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/storage/sqlite"
    "curtaincall.tech/pkg/web"

    "github.com/coreos/go-systemd/activation"
    "github.com/justinas/alice"
    "golang.org/x/crypto/acme/autocert"
    "github.com/bmizerany/pat"


)

func ServeIndex(webRoot string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(webRoot, "index.html"))
	})
}

func main() {

    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    r := retrieving.NewService(s)

    standardMiddleware := alice.New(web.RecoverPanic, web.SecureHeaders, web.CorsHeaders)    
    m := pat.New()

    m.Get("/", http.FileServer(http.Dir("/home/curtaincall/")))
    m.Get("/shows", ServeIndex("/home/curtaincall/"))

    m.Get("/api/shows", http.HandlerFunc(web.RetrieveAllShows(r)))

	handler := standardMiddleware.Then(m)

    listeners, err := activation.Listeners()
    if err != nil {
        log.Fatal(err)
    }
    if len(listeners) != 2 {
        log.Fatal("Missing systemd socket listeners for ports 80 and 443")
    }
    
    httpListener := listeners[0]
    httpsListener := listeners[1]
    
    go http.Serve(httpListener, http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        target := "https://" + r.Host + r.URL.Path
        if len(r.URL.RawQuery) > 0 {
    	    target += "?" + r.URL.RawQuery
        }
        http.Redirect(w, r, target, http.StatusMovedPermanently)
    }))
    
    crt := &autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("curtaincall.tech"),
    }
    
    dir := filepath.Join(os.Getenv("HOME"), ".cache", "golang-autocert")
    if err := os.MkdirAll(dir, 0700); err != nil {
        log.Printf("warning: autocert Manager not using a cache: %v", err)
    } else {
        crt.Cache = autocert.DirCache(dir)
    }
    
    srv := &http.Server{
        Handler:      handler,
        TLSConfig:    crt.TLSConfig(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    //TODO: Zero Downtime Deployment?

	log.Println("Listening on port 443")
	log.Fatal(srv.ServeTLS(httpsListener, "", ""))
}

