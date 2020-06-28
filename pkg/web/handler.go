package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"curtaincall.tech/pkg/retrieving"
	"curtaincall.tech/pkg/storage/sqlite"
)

func RetrieveAllShows(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
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
