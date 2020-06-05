package rest

import (
	"encoding/json"
	"errors"
    "fmt"
    "log"
	"net/http"
	"strconv"

	"curtaincall.tech/pkg/creating"
	"curtaincall.tech/pkg/deleting"
	"curtaincall.tech/pkg/retrieving"
	"curtaincall.tech/pkg/storage/sqlite"

	"github.com/bmizerany/pat"
	"github.com/golang/gddo/httputil/header"
)

func Handler(c creating.Service, r retrieving.Service, d deleting.Service) http.Handler {
	router := pat.New()

	router.Post("/theaters", http.HandlerFunc(createTheater(c)))
	router.Get("/theaters", http.HandlerFunc(retrieveAllTheaters(r)))
	router.Get("/theaters/:id", http.HandlerFunc(retrieveTheater(r)))
	router.Del("/theaters/:id", http.HandlerFunc(deleteTheater(d)))

	router.Get("/theaters/:id/shows", http.HandlerFunc(retrieveAllShows(r)))
	router.Get("/theaters/:id/shows/:showid", http.HandlerFunc(retrieveShow(r)))
	router.Post("/shows", http.HandlerFunc(createShow(c)))

	return router
}

func createTheater(s creating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

        // This chunk checks to see that that Contentp-Type header is present
        // and that it is set to application/json
		if r.Header.Get("Content-Type") != "" {
			value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			if value != "application/json" {
				msg := "Content-Type header is not application/json"
				http.Error(w, msg, http.StatusUnsupportedMediaType)
				return
			}
		}

        // this chunk enforces a maximum 1MB body size
        // Decode() will return "http: request body too large" an HTTP 413
        r.Body = http.MaxBytesReader(w, r.Body, 1048576)

        // decode the body and call DisallowUnknownFields() on it.
        // this will cause Decode() to raise a 
		dec := json.NewDecoder(r.Body)
        dec.DisallowUnknownFields()

		var t creating.Theater
		err := dec.Decode(&t)
		if err != nil {
            var syntaxError *json.SyntaxError
        
            switch {
            case errors.As(err, &syntaxError):
                msg := fmt.Sprintf("Request body contains malformed JSON (at position %d)", syntaxError.Offset)
                http.Error(w, msg, http.StatusBadRequest)

            default:
                log.Println(err.Error())
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            }
            return
		}
		_, err = s.CreateTheater(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New Theater Created!")
	}
}

func createShow(s creating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//var id int
		decoder := json.NewDecoder(r.Body)

		var newShow creating.Show

		err := decoder.Decode(&newShow)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = s.CreateShow(newShow)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New Show Created")
	}
}

func retrieveAllTheaters(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		theaters, err := s.RetrieveAllTheaters()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(theaters)
	}
}

func retrieveTheater(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query().Get(":id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		theater, err := s.RetrieveTheater(id)
		if err != nil {
			if errors.Is(err, sqlite.ErrNoRecord) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(theater)
	}
}

func deleteTheater(s deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query().Get(":id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = s.DeleteTheater(id)
		if err != nil {
			if errors.Is(err, sqlite.ErrNoRecord) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Theater Deleted!")
		w.WriteHeader(http.StatusNoContent)
	}
}

func retrieveAllShows(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		shows, err := s.RetrieveAllShows(id)
		if err != nil {
			if errors.Is(err, sqlite.ErrNoRecord) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shows)
	}
}

func retrieveShow(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query().Get(":showid"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(id)

		show, err := s.RetrieveShow(id)
		if err != nil {
			if errors.Is(err, sqlite.ErrNoRecord) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(show)
	}
}
