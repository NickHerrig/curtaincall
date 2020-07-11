package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/storage/sqlite"
    "curtaincall.tech/pkg/web"

    "github.com/coreos/go-systemd/activation"
    "github.com/justinas/alice"
    "golang.org/x/crypto/acme/autocert"
    "github.com/bmizerany/pat"


)

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func main() {

//    dn, ok := os.LookupEnv("CC_DOMAIN_NAME")
//    if !ok {
//        log.Fatalf("Missing env var CC_DOMAIN_NAME")
//    }

    s, err := sqlite.NewStorage()
    if err != nil {
        log.Fatal(err)
    }

    r := retrieving.NewService(s)

    standardMiddleware := alice.New(web.RecoverPanic, web.SecureHeaders, web.CorsHeaders)    
    m := pat.New()

    m.Get("/api/shows", http.HandlerFunc(web.RetrieveAllShows(r)))

    fileServer := http.FileServer(http.Dir("./frontend/dist/"))
    m.Get("/dist/", http.StripPrefix("/dist", fileServer))

    m.Get("/", http.HandlerFunc(IndexHandler("~/index.html")))

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

