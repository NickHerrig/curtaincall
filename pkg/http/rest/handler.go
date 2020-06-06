package rest

import (
	"encoding/json"
	"errors"
    "fmt"
    "io"
    "log"
	"net/http"
	"strconv"
    "strings"

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

		if r.Header.Get("Content-Type") != "" {
			value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			if value != "application/json" {
				msg := "Content-Type header is not application/json"
				http.Error(w, msg, http.StatusUnsupportedMediaType)
				return
			}
		}

        r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		dec := json.NewDecoder(r.Body)
        dec.DisallowUnknownFields()

		var t creating.Theater
		err := dec.Decode(&t)
		if err != nil {
            var syntaxError *json.SyntaxError
            var unmarshalTypeError *json.UnmarshalTypeError

            switch {
            // Catch any syntax errors in the JSON and send an error message
            // which interpolates the location of the problem to make it
            // easier for the client to fix.
            case errors.As(err, &syntaxError):
                msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
                http.Error(w, msg, http.StatusBadRequest)

            // In some circumstances Decode() may also return an
            // io.ErrUnexpectedEOF error for syntax errors in the JSON. There
            // is an open issue regarding this at
            // https://github.com/golang/go/issues/25956.
            case errors.Is(err, io.ErrUnexpectedEOF):
                msg := fmt.Sprintf("Request body contains badly-formed JSON")
                http.Error(w, msg, http.StatusBadRequest)

            // Catch any type errors, like trying to assign a string in the
            // JSON request body to a int field in our Person struct. We can
            // interpolate the relevant field name and position into the error
            // message to make it easier for the client to fix.
            case errors.As(err, &unmarshalTypeError):
                msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
                http.Error(w, msg, http.StatusBadRequest)

            // Catch the error caused by extra unexpected fields in the request
            // body. We extract the field name from the error message and
            // interpolate it in our custom error message. There is an open
            // issue at https://github.com/golang/go/issues/29035 regarding
            // turning this into a sentinel error.
            case strings.HasPrefix(err.Error(), "json: unknown field "):
                fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
                msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
                http.Error(w, msg, http.StatusBadRequest)

            // An io.EOF error is returned by Decode() if the request body is
            // empty.
            case errors.Is(err, io.EOF):
                msg := "Request body must not be empty"
                http.Error(w, msg, http.StatusBadRequest)

            // Catch the error caused by the request body being too large. Again
            // there is an open issue regarding turning this into a sentinel
            // error at https://github.com/golang/go/issues/30715.
            case err.Error() == "http: request body too large":
                msg := "Request body must not be larger than 1MB"
                http.Error(w, msg, http.StatusRequestEntityTooLarge)

            // Otherwise default to logging the error and sending a 500 Internal
            // Server Error response.
            default:
                log.Println(err.Error())
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            }
            return
        }

    	err = dec.Decode(&struct{}{})
    	if err != io.EOF {
            msg := "Request body must only contain a single JSON object"
            http.Error(w, msg, http.StatusBadRequest)
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
