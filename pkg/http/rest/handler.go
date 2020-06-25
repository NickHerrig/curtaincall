package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"curtaincall.tech/pkg/creating"
	"curtaincall.tech/pkg/deleting"
	"curtaincall.tech/pkg/retrieving"
	"curtaincall.tech/pkg/storage/sqlite"

	"github.com/bmizerany/pat"
    "github.com/justinas/alice"
)

func Handler(c creating.Service, r retrieving.Service, d deleting.Service) http.Handler {
    standardMiddleware := alice.New(secureHeaders, corsHeaders)    

	router := pat.New()

	router.Get("/shows", http.HandlerFunc(retrieveAllShows(r)))

	return standardMiddleware.Then(router)
}

func retrieveAllShows(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		shows, err := s.RetrieveAllShows()
		if err != nil {
			if errors.Is(err, sqlite.ErrNoRecord) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shows)
	}
}

