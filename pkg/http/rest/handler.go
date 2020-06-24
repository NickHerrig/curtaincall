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
)

func Handler(c creating.Service, r retrieving.Service, d deleting.Service) http.Handler {
	router := pat.New()

	router.Get("/shows", http.HandlerFunc(retrieveAllShows(r)))

	return router
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

		//SET HEADERS IN MIDDLEWARE
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shows)
	}
}

